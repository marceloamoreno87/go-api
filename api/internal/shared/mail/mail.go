package mail

import (
	"github.com/marceloamoreno/goapi/config"
	sendgrid "github.com/marceloamoreno/goapi/internal/shared/mail/send_grid"
	"github.com/marceloamoreno/goapi/internal/shared/mail/smtp"
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
	switch config.Environment.GetMailDriver() {
	case "smtp":
		return smtp.NewMail()
	case "sendgrid":
		return sendgrid.NewSendGrid()
	default:
		return smtp.NewMail()
	}
}
