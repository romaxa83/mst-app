package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"github.com/romaxa83/mst-app/gin-app/internal/services"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.UserSignUp)
		users.POST("/sign-in", h.userSignIn)
		users.POST("/auth/refresh", h.userRefresh)
	}
}

type userSignInInput struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

// @Summary User SignUp
// @Tags users-auth
// @Description create user account
// @ID user-sign-up
// @Accept  json
// @Produce  json
// @Param input body userSignUpInput true "sign up info"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/sign-up [post]
func (h *Handler) UserSignUp(c *gin.Context) {
	var inp services.UserSignUpInput

	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.Services.Users.SignUp(c.Request.Context(), services.UserSignUpInput{
		Name:     inp.Name,
		Email:    inp.Email,
		Phone:    inp.Phone,
		Password: inp.Password,
	})
	if err != nil {
		if errors.Is(err, domains.ErrUserAlreadyExists) {
			errorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

// @Summary User SignIn
// @Tags users-auth
// @Description user sign in
// @ModuleID userSignIn
// @Accept  json
// @Produce  json
// @Param input body userSignInInput true "sign up info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/sign-in [post]
func (h *Handler) userSignIn(c *gin.Context) {
	var inp userSignInInput
	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.Services.Users.SignIn(c.Request.Context(), services.UserSignInInput{
		Email:    inp.Email,
		Password: inp.Password,
	})
	if err != nil {
		if errors.Is(err, domains.ErrUserNotFound) {
			errorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

type refreshInput struct {
	Token string `json:"token" binding:"required"`
}

// @Summary User Refresh Tokens
// @Tags users-auth
// @Description user refresh tokens
// @Accept  json
// @Produce  json
// @Param input body refreshInput true "sign up info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/auth/refresh [post]
func (h *Handler) userRefresh(c *gin.Context) {
	var inp refreshInput
	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.Services.Users.RefreshTokens(c.Request.Context(), inp.Token)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}
