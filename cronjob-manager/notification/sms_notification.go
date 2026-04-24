package notification

import (
	"context"
	"fmt"
	"log/slog"
	"regexp"
	"time"
)

type SmsNotification struct {
	logger *slog.Logger
}

func NewSmsNotification(logger *slog.Logger) *SmsNotification {
	return &SmsNotification{logger: logger}
}


func (sms *SmsNotification) Send(ctx context.Context, phone string, text string) error {
	cleanNumber := sms.sanitizeNumber(phone)
	
	// logic here: api call, response, ..
    fmt.Printf("\nSending message '%s' to '%s'\n", text, cleanNumber)
	// .....

	sms.logger.Info("Message sent", 
		"phone", cleanNumber, 
		"sent_at", time.Now(),
	)
	return nil
}

func (p *SmsNotification) sanitizeNumber(phone string) string {
	reg := regexp.MustCompile(`[^0-9]`)
	return reg.ReplaceAllString(phone, "")
}