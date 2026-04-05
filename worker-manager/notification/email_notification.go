package notification

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

type EmailNotification struct {
	logger *slog.Logger
}

func NewEmailNotification(logger *slog.Logger) *EmailNotification {
	return &EmailNotification{logger: logger}
}

func (em *EmailNotification) Send(ctx context.Context, from string, to string, data interface{}) error {
	fromEmail := em.sanitizeEmail(from)
	toEmail := em.sanitizeEmail(to)
	
	// logic here: template, connection, sending....
    fmt.Printf("Sending email from '%s' to '%s", fromEmail, toEmail)
	// .....

	em.logger.Info("Message sent", 
		"from", fromEmail, 
		"to", toEmail,
		"sent_at", time.Now(),
	)
	return nil
}

func (em *EmailNotification) sanitizeEmail(rawEmail string) string {
	email := rawEmail
	return  email
}