package config

import (
	"github.com/ortizdavid/go-nopain/mailer"
)

func DefaultEmailService() *mailer.EmailService {
	return &mailer.EmailService{
		SMTPHost: MailSMTPHost(),
		SMTPPort: MailSMTPPort(),
		Username: MailUser(),
		Password: MailPassword(),
	}
}

func MailUser() string {
	return GetEnv("MAIL_USER")
}

func MailPassword() string {
	return GetEnv("MAIL_PASSWORD")
}

func MailSMTPHost() string {
	return GetEnv("MAIL_SMTP_HOST")
}

func MailSMTPPort() string {
	return GetEnv("MAIL_SMTP_PORT")
}
