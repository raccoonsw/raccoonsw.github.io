package grpc_email_server

import (
	"context"
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
)

type Server struct {
	SmtpPort     int
	SmtpHost     string
	UserEmail    string
	UserPassword string
}

func (s *Server) Send(ctx context.Context, in *Request) (*Response, error) {
	msg := s.CreateEmail(in)
	err := s.SendEmailSmtp(msg)
	if err != nil {
		return nil, err
	}
	return &Response{Response: "Email sent."}, nil
}

func (s *Server) CreateEmail(in *Request) *gomail.Message {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "from@gmail.com")
	msg.SetHeader("To", in.Email)
	msg.SetHeader("Subject", fmt.Sprintf("Order №%d created!", in.OrderId))
	// Set E-Mail body. You can set plain text or html with text/html
	msg.SetBody("text/plain", fmt.Sprintf("You ordered the item №%d from our catalog. If it was not you, contact support team.", in.ItemId))
	return msg
}

func (s *Server) SendEmailSmtp(m *gomail.Message) (err error) {
	d := gomail.NewDialer(s.SmtpHost, s.SmtpPort, s.UserEmail, s.UserPassword)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
