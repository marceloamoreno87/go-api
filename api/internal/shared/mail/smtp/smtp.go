package smtp

// TODO: Refactor
import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"

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
	return &Mail{
		From: config.Environment.GetMailFrom(),
	}
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

func (m *Mail) SetSubject(subject string) {
	m.Subject = subject
}

// TODO: Refactor
func (m *Mail) SetBody(filename string, data any) {
	t, err := template.ParseFiles(filename + ".html")
	if err != nil {
		log.Println(err)
	}
	var tpl bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	tpl.Write([]byte(fmt.Sprintf("From: %s \nSubject: %s \n%s\n\n", m.From, m.Subject, mimeHeaders)))
	if err := t.Execute(&tpl, data); err != nil {
		log.Println(err)
	}
	result := tpl.String()
	m.Body = result
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
