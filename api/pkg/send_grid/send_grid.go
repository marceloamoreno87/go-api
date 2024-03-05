package sendgrid

// TODO: Refactor
import (
	"bytes"
	"log"
	"log/slog"
	"text/template"

	"github.com/marceloamoreno/goapi/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGrid struct {
	Name        string
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
		Name: config.Environment.GetMailName(),
	}
}

func (m *SendGrid) Send() (err error) {
	from := mail.NewEmail("<"+m.Name+"> ", m.From)
	subject := m.Subject
	for _, to := range m.To {
		to := mail.NewEmail(to, to)
		message := mail.NewSingleEmail(from, subject, to, "", m.Body)
		client := sendgrid.NewSendClient(config.Environment.GetSendgridApiKey())
		_, err = client.Send(message)
		if err != nil {
			slog.Info(err.Error())
		}
	}
	return err
}

func (m *SendGrid) SetTo(to []string) {
	m.To = to
}

func (m *SendGrid) SetSubject(subject string) {
	m.Subject = subject
}

func (m *SendGrid) SetBody(filename string, data any) {
	t, err := template.ParseFiles("internal/views/" + filename + ".html")
	if err != nil {
		log.Println(err)
	}
	var tpl bytes.Buffer
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
