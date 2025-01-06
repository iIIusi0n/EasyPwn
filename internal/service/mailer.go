package service

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"net/smtp"

	pb "easypwn/internal/api"
	"easypwn/internal/data"
)

type MailerServiceConfig struct {
	SmtpHost    string
	SmtpTlsPort string
	SmtpUser    string
	SmtpPass    string
}

type MailerService struct {
	config MailerServiceConfig

	pb.UnimplementedMailerServer
}

func NewMailerService(ctx context.Context, config MailerServiceConfig) *MailerService {
	return &MailerService{
		config: config,
	}
}

func (s *MailerService) SendConfirmationEmail(ctx context.Context, req *pb.SendConfirmationEmailRequest) (*pb.SendConfirmationEmailResponse, error) {
	exists, err := s.hasRecentConfirmation(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check recent confirmations: %v", err)
	}
	if exists {
		return nil, fmt.Errorf("confirmation code already sent within last 3 minutes")
	}

	code := generateConfirmationCode()

	if err := s.storeConfirmationCode(ctx, req.Email, code); err != nil {
		return nil, fmt.Errorf("failed to store confirmation code: %v", err)
	}

	if err := s.sendEmail(req.Email, code); err != nil {
		return nil, fmt.Errorf("failed to send confirmation email: %v", err)
	}

	return &pb.SendConfirmationEmailResponse{
		Code: code,
	}, nil
}

func (s *MailerService) storeConfirmationCode(ctx context.Context, email, code string) error {
	db := data.GetDB()

	_, err := db.ExecContext(ctx,
		"INSERT INTO email_confirmation (id, email, code) VALUES (UUID_TO_BIN(UUID()), ?, ?)",
		email, code)

	return err
}

func (s *MailerService) sendEmail(to, code string) error {
	auth := smtp.PlainAuth("", s.config.SmtpUser, s.config.SmtpPass, s.config.SmtpHost)

	msg := []byte(fmt.Sprintf(`From: EasyPwn <%s>
To: %s
Subject: Email Confirmation

Your confirmation code is: %s

Best regards,
EasyPwn Team
`, s.config.SmtpUser, to, code))

	addr := fmt.Sprintf("%s:%s", s.config.SmtpHost, s.config.SmtpTlsPort)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         s.config.SmtpHost,
	}

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()

	if err = c.StartTLS(tlsconfig); err != nil {
		return err
	}

	if err = c.Auth(auth); err != nil {
		return err
	}

	if err = c.Mail(s.config.SmtpUser); err != nil {
		return err
	}
	if err = c.Rcpt(to); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write(msg)
	return err
}

func generateConfirmationCode() string {
	bytes := make([]byte, 3)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (s *MailerService) hasRecentConfirmation(ctx context.Context, email string) (bool, error) {
	db := data.GetDB()

	var count int
	err := db.QueryRowContext(ctx, `
		SELECT COUNT(*) 
		FROM email_confirmation 
		WHERE email = ? 
		AND created_at > DATE_SUB(NOW(), INTERVAL 3 MINUTE)`,
		email).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *MailerService) GetConfirmationCode(ctx context.Context, req *pb.GetConfirmationCodeRequest) (*pb.GetConfirmationCodeResponse, error) {
	db := data.GetDB()

	var code string
	err := db.QueryRowContext(ctx, `
		SELECT code 
		FROM email_confirmation 
		WHERE email = ?
		AND created_at > DATE_SUB(NOW(), INTERVAL 3 MINUTE)`,
		req.Email).Scan(&code)

	if err != nil {
		return nil, err
	}

	return &pb.GetConfirmationCodeResponse{
		Code: code,
	}, nil
}
