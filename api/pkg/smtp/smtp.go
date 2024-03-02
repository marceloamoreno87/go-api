package smtp

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/marceloamoreno/goapi/config"
)

type Mail struct {
	To          []string
	From        string
	Subject     string
	Body        string
	CC          string
	BCC         string
	Attachments []string
}

func NewMail() *Mail {
	return &Mail{}
}

func (m *Mail) Send() (err error) {
	err = smtp.SendMail(fmt.Sprintf("%s:%s", config.Environment.GetMailHost(), config.Environment.GetMailPort()), nil, m.From, m.To, []byte(m.Body))
	if err != nil {
		log.Println(err)
		return err
	}
	return
}

func (m *Mail) SetTo(to []string) {
	m.To = to
}

func (m *Mail) SetFrom(from string) {
	m.From = from
}

func (m *Mail) SetSubject(subject string) {
	m.Subject = subject
}

func (m *Mail) SetBody(body string) {
	m.Body = body
}

func (m *Mail) SetCC(cc string) {
	m.CC = cc
}

func (m *Mail) SetBCC(bcc string) {
	m.BCC = bcc
}

func (m *Mail) SetAttachments(attachments []string) {
	m.Attachments = attachments
}
