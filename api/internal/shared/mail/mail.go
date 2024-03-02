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
	SetFrom(from string)
	SetSubject(subject string)
	SetBody(body string)
	SetCC(cc string)
	SetBCC(bcc string)
	SetAttachments(attachments []string)
}

func NewMail() MailInterface {
	switch {
	case config.Environment.GetMailDriver() == "smtp":
		return smtp.NewMail()
	case config.Environment.GetMailDriver() == "mailersend":
		return mailersend.NewMailerSend()
	case config.Environment.GetMailDriver() == "sendgrid":
		return sendgrid.NewSendGrid()
	default:
		return nil
	}
}
