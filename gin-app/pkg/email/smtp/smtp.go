package smtp

import (
	"github.com/go-gomail/gomail"
	"github.com/pkg/errors"
	"github.com/romaxa83/mst-app/gin-app/pkg/email"
	"github.com/romaxa83/mst-app/gin-app/pkg/logger"
)

type SMTPSender struct {
	from string
	pass string
	host string
	port int
}

func NewSMTPSender(from, pass, host string, port int) (*SMTPSender, error) {
	if !email.IsEmailValid(from) {
		return nil, errors.New("invalid from email")
	}

	return &SMTPSender{from: from, pass: pass, host: host, port: port}, nil
}

func (s *SMTPSender) Send(input email.SendEmailInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", input.To)
	msg.SetHeader("Subject", input.Subject)
	msg.SetBody("text/html", input.Body)

	//d := gomail.NewDialer(s.host, s.port, s.from, s.pass)

	dialer := gomail.Dialer{
		Host: s.host,
		Port: s.port,
	}
	if err := dialer.DialAndSend(msg); err != nil {
		return errors.Wrap(err, "failed to sent email via smtp")
	}

	logger.Warn("SEND EMAIL")

	return nil
}
