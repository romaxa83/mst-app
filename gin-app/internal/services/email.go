package services

import (
	"fmt"
	"github.com/romaxa83/mst-app/gin-app/internal/config"
	emailProvider "github.com/romaxa83/mst-app/gin-app/pkg/email"
)

const (
	verificationLinkTmpl = "https://%s/verification?code=%s" // https://<school host>/verification?code=<verification_code>
)

type EmailService struct {
	config config.EmailConfig
	sender emailProvider.Sender
}

// Structures used for templates.
type verificationEmailInput struct {
	VerificationLink string
}

func NewEmailsService(config config.EmailConfig, sender emailProvider.Sender) *EmailService {
	return &EmailService{
		config: config,
		sender: sender,
	}
}

func (s EmailService) SendVerificationEmail(input VerificationEmailInput) error {
	subject := fmt.Sprintf(s.config.Subjects.Verification, input.Name)
	// todo.txt вынести domain в настройки
	templateInput := verificationEmailInput{s.createVerificationLink("127.0.0.1", input.VerificationCode)}
	sendInput := emailProvider.SendEmailInput{Subject: subject, To: input.Email}

	if err := sendInput.GenerateBodyFromHTML(s.config.Templates.Verification, templateInput); err != nil {
		return err
	}

	return s.sender.Send(sendInput)
}

func (s *EmailService) createVerificationLink(domain, code string) string {
	return fmt.Sprintf(verificationLinkTmpl, domain, code)
}
