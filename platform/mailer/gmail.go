package mailer

import (
	"fmt"
	"github.com/spf13/viper"
	"net/smtp"
)

type gmailMailer struct {
	fromEmail         string
	appPass           string
	smtpServerAddress string
}

func (g *gmailMailer) SendConfirmationCode(toEmail string, code string) error {
	confirmationUrl := fmt.Sprintf("%s/%s", viper.GetString("confirm_url"), code)
	msg := fmt.Sprintf("Please click the following link to confirm your account: %s", confirmationUrl)
	return g.Send(&EmailMessage{
		Subject:  "Account confirmation",
		ToEmail:  toEmail,
		TextBody: msg,
	})
}

func (g *gmailMailer) SendPasswordResetCode(toEmail, code string) error {
	//TODO implement me
	panic("implement me")
}

func (g *gmailMailer) Send(message *EmailMessage) error {
	fmt.Println("[Gmail Mailer] start")

	err := smtp.SendMail(
		fmt.Sprintf("%s:587", g.smtpServerAddress),
		smtp.PlainAuth("", "sultan.imanhodjaev@gmail.com", g.appPass, g.smtpServerAddress),
		g.fromEmail,
		[]string{message.ToEmail},
		[]byte(message.TextBody),
	)

	if err != nil {
		return err
	}

	fmt.Println("[Gmail Mailer] end")
	return nil
}

func NewGmailMailer() Mailer {
	return &gmailMailer{
		fromEmail:         viper.GetString("from_email"),
		appPass:           viper.GetString("gmail_app_pass"),
		smtpServerAddress: "smtp.gmail.com",
	}
}
