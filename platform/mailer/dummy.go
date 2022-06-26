package mailer

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type dummyMailer struct{}

func (d *dummyMailer) SendConfirmationCode(toEmail string, code string) error {
	link := fmt.Sprintf("https://%s/confirm/%s", viper.GetString("app_host"), code)
	return d.Send(&EmailMessage{
		Subject:  "Your confirmation link",
		ToEmail:  toEmail,
		TextBody: fmt.Sprintf("Please use the following confirmation link: %s", link),
		HTMLBody: fmt.Sprintf("Please use the following confirmation link: %s", link),
	})
}

func (d *dummyMailer) SendPasswordResetCode(toEmail string, code string) error {
	link := fmt.Sprintf("https://%s/reset-password/%s", viper.GetString("app_host"), code)
	return d.Send(&EmailMessage{
		Subject:  "Your password reset link",
		ToEmail:  toEmail,
		TextBody: fmt.Sprintf("Please use the following link to reset your password: %s", link),
		HTMLBody: fmt.Sprintf("Please use the following link to reset your password: %s", link),
	})
}

func NewDummyMailer() Mailer {
	return &dummyMailer{}
}

func (d *dummyMailer) Send(message *EmailMessage) error {
	fmt.Println("[Dummy Mailer] start")
	spew.Dump(message)
	fmt.Println("[Dummy Mailer] end")
	return nil
}
