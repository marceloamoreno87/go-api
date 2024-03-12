package event

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/auth/entity"
	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

type PasswordForgotEmailEventInterface interface {
	Send()
}

type PasswordForgotEmailEvent struct {
	UserValidation entity.UserValidationInterface
	Mail           mail.MailInterface
}

func NewPasswordForgotEmailEvent(userValidation entity.UserValidationInterface) *PasswordForgotEmailEvent {
	return &PasswordForgotEmailEvent{
		UserValidation: userValidation,
		Mail:           mail.NewMail(),
	}
}

func (e *PasswordForgotEmailEvent) Send() {
	e.Mail.SetTo([]string{e.UserValidation.GetUser().GetEmail()})
	e.Mail.SetSubject("Recuperação de senha!")
	e.Mail.SetBody("forgot_password", e.UserValidation)
	err := e.Mail.Send()
	if err != nil {
		slog.Error("error sending email", err)
	}
}
