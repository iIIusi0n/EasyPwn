package service

import (
	"context"
	"os"
	"testing"

	pb "easypwn/internal/api"
)

func getMailerConfig() MailerServiceConfig {
	return MailerServiceConfig{
		SmtpHost:    os.Getenv("MAILER_SMTP_HOST"),
		SmtpTlsPort: os.Getenv("MAILER_SMTP_TLS_PORT"),
		SmtpUser:    os.Getenv("MAILER_SMTP_USER"),
		SmtpPass:    os.Getenv("MAILER_SMTP_PASS"),
	}
}

func TestSendEmailConfirmation(t *testing.T) {
	ctx := context.Background()
	config := getMailerConfig()
	mailerService := NewMailerService(ctx, config)

	t.Run("SuccessfulSend", func(t *testing.T) {
		resp, err := mailerService.SendEmailConfirmation(ctx, &pb.SendEmailConfirmationRequest{
			Email: "test@example.com",
		})

		if err != nil {
			t.Errorf("SendEmailConfirmation() error = %v", err)
			return
		}

		if resp.Code == "" {
			t.Error("SendEmailConfirmation() returned empty confirmation code")
		}
	})

	t.Run("DuplicateRequest", func(t *testing.T) {
		_, err := mailerService.SendEmailConfirmation(ctx, &pb.SendEmailConfirmationRequest{
			Email: "duplicate@example.com",
		})
		if err != nil {
			t.Errorf("First SendEmailConfirmation() error = %v", err)
			return
		}

		_, err = mailerService.SendEmailConfirmation(ctx, &pb.SendEmailConfirmationRequest{
			Email: "duplicate@example.com",
		})
		if err == nil {
			t.Error("Expected error for duplicate request within 3 minutes, got nil")
		}
	})

	t.Run("InvalidEmail", func(t *testing.T) {
		_, err := mailerService.SendEmailConfirmation(ctx, &pb.SendEmailConfirmationRequest{
			Email: "",
		})
		if err == nil {
			t.Error("Expected error for invalid email, got nil")
		}
	})
}

func TestGenerateConfirmationCode(t *testing.T) {
	code1 := generateConfirmationCode()
	code2 := generateConfirmationCode()

	if code1 == "" {
		t.Error("generateConfirmationCode() returned empty string")
	}

	if code1 == code2 {
		t.Error("generateConfirmationCode() returned same code twice")
	}

	if len(code1) != 6 {
		t.Errorf("generateConfirmationCode() returned code of length %d, want 6", len(code1))
	}
}
