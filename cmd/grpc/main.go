package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	lib "google.golang.org/grpc"
	"log"
	"net"
	email "restApiProject/grpc_email_server"
)

type Specification struct {
	SmtpPort     int    `required:"true"`
	SmtpHost     string `required:"true"`
	Port         string `required:"true"`
	Host         string `required:"true"`
	UserEmail    string `required:"true"`
	UserPassword string `required:"true"`
}

func main() {
	var s Specification
	err := envconfig.Process("grpc", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.Host, s.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := email.Server{SmtpHost: s.SmtpHost, SmtpPort: s.SmtpPort, UserEmail: s.UserEmail, UserPassword: s.UserPassword}

	grpcServer := lib.NewServer()

	email.RegisterEmailServiceServer(grpcServer, &srv)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
