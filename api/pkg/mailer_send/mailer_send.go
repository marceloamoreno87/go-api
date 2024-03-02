package mailersend

import (
	"context"

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
	return &MailerSend{}
}

func (m *MailerSend) Send() (err error) {
	ctx := context.Background()
	ms := mailersend.NewMailersend(config.Environment.GetMailerSendApiKey())

	from := mailersend.From{
		Name:  m.From,
		Email: m.From,
	}

	// TODO: REFACTOR
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

func (m *MailerSend) SetFrom(from string) {
	m.From = from
}

func (m *MailerSend) SetSubject(subject string) {
	m.Subject = subject
}

func (m *MailerSend) SetBody(body string) {
	m.Body = body
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
