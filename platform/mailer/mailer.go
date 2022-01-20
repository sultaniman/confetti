package mailer

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
