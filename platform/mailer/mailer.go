package mailer

import (
	"github.com/spf13/viper"
)

type EmailMessage struct {
	Subject  string
	ToEmail  string
	TextBody string
	HTMLBody string
}

type Mailer interface {
	Send(message *EmailMessage) error
	SendConfirmationCode(toEmail string, code string) error
	SendPasswordResetCode(toEmail, code string) error
}

func GetMailer() Mailer {
	switch viper.GetString("mailer") {
	case "gmail":
		return NewGmailMailer()
	case "mailjet":
		return NewMJMailer()
	default:
		return NewDummyMailer()
	}
}
