package services

import (
	"context"
	"fmt"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"github.com/romaxa83/mst-app/gin-app/internal/repositories"
	"github.com/romaxa83/mst-app/gin-app/pkg/auth"
	"github.com/romaxa83/mst-app/gin-app/pkg/hash"
	"github.com/romaxa83/mst-app/gin-app/pkg/otp"
	"time"
)

type UsersService struct {
	repo repositories.Users

	hasher                 hash.PasswordHasher
	tokenManager           auth.TokenManager
	accessTokenTTL         time.Duration
	refreshTokenTTL        time.Duration
	verificationCodeLength int

	otpGenerator otp.Generator

	emailService Emails
}

func NewUsersService(
	repo repositories.Users,
	hasher hash.PasswordHasher,
	tokenManager auth.TokenManager,
	accessTTL,
	refreshTTL time.Duration,
	emailService *EmailService,
	otpGenerator otp.Generator,
	verificationCodeLength int,
) *UsersService {
	return &UsersService{
		repo:                   repo,
		hasher:                 hasher,
		tokenManager:           tokenManager,
		accessTokenTTL:         accessTTL,
		refreshTokenTTL:        refreshTTL,
		emailService:           emailService,
		otpGenerator:           otpGenerator,
		verificationCodeLength: verificationCodeLength,
	}
}

func (s *UsersService) SignUp(ctx context.Context, input UserSignUpInput) (int, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return 0, err
	}

	verificationCode := s.otpGenerator.RandomSecret(s.verificationCodeLength)

	user := domains.User{
		Name:      input.Name,
		Password:  passwordHash,
		Phone:     input.Phone,
		Email:     input.Email,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.emailService.SendVerificationEmail(VerificationEmailInput{
		Email:            user.Email,
		Name:             user.Name,
		VerificationCode: verificationCode,
	})

	return s.repo.Create(ctx, user)
}

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(input.Password)

	if err != nil {
		return Tokens{}, err
	}

	user, err := s.repo.GetByCredentials(ctx, input.Email, passwordHash)
	if err != nil {
		return Tokens{}, err
	}

	return s.createSession(ctx, user.ID)
}

func (s *UsersService) RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error) {
	user, err := s.repo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return Tokens{}, err
	}

	return s.createSession(ctx, user.ID)
}

func (s *UsersService) createSession(ctx context.Context, userId int) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(fmt.Sprintf("%v", userId), s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}

	session := domains.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}

	err = s.repo.SetSession(ctx, userId, session)

	return res, err
}
