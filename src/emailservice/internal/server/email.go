package server

import (
	"context"
	. "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	. "github.com/golang/protobuf/ptypes/empty"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"strings"
)

func NewEmailServiceServer() *emailServiceServer {
	log.Printf("Registered emailServiceServer handler")
	return &emailServiceServer{}
}

type emailServiceServer struct{}


func (s *emailServiceServer) Send(ctx context.Context, req *SendEmailRequest) (*Empty, error) {
	from := mail.NewEmail(req.Sender, req.SenderEmail)
	to := mail.NewEmail(req.Recipient, req.RecipientEmail)

	subject := req.Subject

	plainTextContent := req.Content
	htmlContent := strings.Replace(req.Content,"\n","<br>",-1)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)

	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, status.Error(codes.Internal, response.Body)
	}

	return &Empty{}, nil
}
