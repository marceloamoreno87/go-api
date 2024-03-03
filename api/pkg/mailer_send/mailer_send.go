package mailersend

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"text/template"

	"github.com/mailersend/mailersend-go"
	"github.com/marceloamoreno/goapi/config"
)

type MailerSend struct {
	To          []string
	From        string
	Subject     string
	Body        string
	CC          string
	BCC         string
	Attachments []string
}

func NewMailerSend() *MailerSend {
	return &MailerSend{
		From: config.Environment.GetMailFrom(),
	}
}

// TODO: REFACTOR
func (m *MailerSend) Send() (err error) {
	ctx := context.Background()
	ms := mailersend.NewMailersend(config.Environment.GetMailerSendApiKey())

	from := mailersend.From{
		Name:  m.From,
		Email: m.From,
	}

	recipients := []mailersend.Recipient{
		{
			Name:  m.To[0],
			Email: m.To[0],
		},
	}
	message := ms.Email.NewMessage()
	message.SetFrom(from)
	message.SetSubject(m.Subject)
	message.SetHTML(m.Body)
	message.SetRecipients(recipients)
	_, err = ms.Email.Send(ctx, message)
	return err
}

func (m *MailerSend) SetTo(to []string) {
	m.To = to
}

func (m *MailerSend) SetSubject(subject string) {
	m.Subject = subject
}

// TODO: Refactor
func (m *MailerSend) SetBody(filename string, data any) {
	t, err := template.ParseFiles("internal/views/" + filename + ".html")
	if err != nil {
		log.Println(err)
	}
	var tpl bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	tpl.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", m.Subject, mimeHeaders)))
	if err := t.Execute(&tpl, data); err != nil {
		log.Println(err)
	}
	result := tpl.String()
	m.Body = result
}

func (m *MailerSend) SetCC(cc string) {
	m.CC = cc
}

func (m *MailerSend) SetBCC(bcc string) {
	m.BCC = bcc
}

func (m *MailerSend) SetAttachments(attachments []string) {
	m.Attachments = attachments
}
