package sendgrid

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

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
	return &SendGrid{
		From: config.Environment.GetMailFrom(),
	}
}

// TODO: REFACTOR
func (m *SendGrid) Send() (err error) {
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

func (m *SendGrid) SetSubject(subject string) {
	m.Subject = subject
}

// TODO: Refactor
func (m *SendGrid) SetBody(filename string, data any) {
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

func (m *SendGrid) SetCC(cc string) {
	m.CC = cc
}

func (m *SendGrid) SetBCC(bcc string) {
	m.BCC = bcc
}

func (m *SendGrid) SetAttachments(attachments []string) {
	m.Attachments = attachments
}
