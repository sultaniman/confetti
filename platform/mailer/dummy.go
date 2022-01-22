package mailer

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type dummyMailer struct{}

func NewDummyMailer() Mailer {
	return &dummyMailer{}
}

func (d *dummyMailer) Send(message *EmailMessage) error {
	fmt.Println("[Dummy Mailer] start")
	spew.Dump(message)
	fmt.Println("[Dummy Mailer] end")
	return nil
}
