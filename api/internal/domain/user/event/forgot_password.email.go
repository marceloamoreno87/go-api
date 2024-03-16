package event

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

type PasswordForgotEmailEventInterface interface {
	Send()
}

type PasswordForgotEmailEventInputDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Hash  string `json:"hash"`
}

type PasswordForgotEmailEvent struct {
	Mail mail.MailInterface
	Data PasswordForgotEmailEventInputDTO
}

func NewPasswordForgotEmailEvent(data PasswordForgotEmailEventInputDTO) *PasswordForgotEmailEvent {
	return &PasswordForgotEmailEvent{
		Data: data,
		Mail: mail.NewMail(),
	}
}

func (e *PasswordForgotEmailEvent) Send() {
	e.Mail.SetTo([]string{e.Data.Email})
	e.Mail.SetSubject("Recuperação de senha!")
	e.Mail.SetBody("forgot_password", e.Data)
	err := e.Mail.Send()
	if err != nil {
		slog.Error("error sending email", err)
	}
}
