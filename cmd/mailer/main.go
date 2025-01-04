package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "easypwn/internal/api"
	"easypwn/internal/service"

	"google.golang.org/grpc"
)

var (
	listenPort  = os.Getenv("MAILER_LISTEN_PORT")
	smtpHost    = os.Getenv("MAILER_SMTP_HOST")
	smtpTlsPort = os.Getenv("MAILER_SMTP_TLS_PORT")
	smtpUser    = os.Getenv("MAILER_SMTP_USER")
	smtpPass    = os.Getenv("MAILER_SMTP_PASS")
)

func init() {
	if listenPort == "" {
		listenPort = "50053"
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMailerServer(s, service.NewMailerService(context.Background(), service.MailerServiceConfig{
		SmtpHost:    smtpHost,
		SmtpTlsPort: smtpTlsPort,
		SmtpUser:    smtpUser,
		SmtpPass:    smtpPass,
	}))

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
