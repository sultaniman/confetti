package mailer

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type sendgridMailer struct{}

func (s *sendgridMailer) SendConfirmationCode(toEmail string, code string) error {
	//TODO implement me
	panic("implement me")
}

func (s *sendgridMailer) SendPasswordResetCode(toEmail, code string) error {
	//TODO implement me
	panic("implement me")
}

func NewSendgridMailer() Mailer {
	return &sendgridMailer{}
}

func (s *sendgridMailer) Send(message *EmailMessage) error {
	fmt.Println("[Dummy Mailer] start")
	spew.Dump(message)
	fmt.Println("[Dummy Mailer] end")
	return nil
}
