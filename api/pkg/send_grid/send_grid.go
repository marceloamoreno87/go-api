package sendgrid

import (
	"github.com/marceloamoreno/goapi/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
	To          []string
	From        string
	Subject     string
	Body        string
	CC          string
	BCC         string
	Attachments []string
}

func NewSendGrid() *SendGrid {
	return &SendGrid{}
}

func (m *SendGrid) Send() (err error) {

	// TODO: REFACTOR
	from := mail.NewEmail("Example User", "test@example.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "test@example.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.Environment.GetSendgridApiKey())
	_, err = client.Send(message)

	return err
}

func (m *SendGrid) SetTo(to []string) {
	m.To = to
}

func (m *SendGrid) SetFrom(from string) {
	m.From = from
}

func (m *SendGrid) SetSubject(subject string) {
	m.Subject = subject
}

func (m *SendGrid) SetBody(body string) {
	m.Body = body
}

func (m *SendGrid) SetCC(cc string) {
	m.CC = cc
}

func (m *SendGrid) SetBCC(bcc string) {
	m.BCC = bcc
}

func (m *SendGrid) SetAttachments(attachments []string) {
	m.Attachments = attachments
}
