package mail

import (
	"github.com/marceloamoreno/goapi/config"
	mailersend "github.com/marceloamoreno/goapi/pkg/mailer_send"
	sendgrid "github.com/marceloamoreno/goapi/pkg/send_grid"
	"github.com/marceloamoreno/goapi/pkg/smtp"
)

type MailInterface interface {
	Send() error
	SetTo(to []string)
	SetSubject(subject string)
	SetBody(filename string, data any)
	SetCC(cc string)
	SetBCC(bcc string)
	SetAttachments(attachments []string)
}

func NewMail() MailInterface {
	if config.Environment.GetEnv() == "development" {
		return smtp.NewMail()
	}

	switch config.Environment.GetMailDriver() {
	case "smtp":
		return smtp.NewMail()
	case "mailersend":
		return mailersend.NewMailerSend()
	case "sendgrid":
		return sendgrid.NewSendGrid()
	default:
		return smtp.NewMail()
	}
}
