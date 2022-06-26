package mailer

import (
	"fmt"
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type mjMailer struct {
	fromEmail string
	apiKey    string
	apiSecret string
}

func (g *mjMailer) SendConfirmationCode(toEmail string, code string) error {
	confirmationUrl := fmt.Sprintf("%s/%s", viper.GetString("confirm_url"), code)
	msg := fmt.Sprintf("Please click the following link to confirm your account: %s", confirmationUrl)
	return g.Send(&EmailMessage{
		Subject:  "Account confirmation",
		ToEmail:  toEmail,
		TextBody: msg,
	})
}

func (g *mjMailer) SendPasswordResetCode(toEmail, code string) error {
	//TODO implement me
	panic("implement me")
}

func (g *mjMailer) Send(message *EmailMessage) error {
	log.Info().Msg("[MJ] Sending message start")
	mailjetClient := mailjet.NewMailjetClient(g.apiKey, g.apiSecret)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: g.fromEmail,
				Name:  "Confetti",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: message.ToEmail,
				},
			},
			Subject:  message.Subject,
			TextPart: message.TextBody,
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.
			Err(err).
			Msg("Unable to send mail using Mailjet")
		return err
	}

	log.Info().Msg("[MJ] Message sent")
	return nil
}

func NewMJMailer() Mailer {
	return &mjMailer{
		fromEmail: viper.GetString("from_email"),
		apiKey:    viper.GetString("mj_key"),
		apiSecret: viper.GetString("mj_secret"),
	}
}
