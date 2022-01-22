package mailer

import (
	"github.com/spf13/viper"
)

type EmailMessage struct {
	Subject   string
	ToEmail   string
	FromEmail string
	TextBody  string
	HTMLBody  string
}

type Mailer interface {
	Send(message *EmailMessage) error
}

func GetMailer() Mailer {
	switch viper.GetString("mailer") {
	case "sendgrid":
		return NewSendgridMailer()
	default:
		return NewDummyMailer()
	}
}
