package event

import (
	"log/slog"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

type PasswordForgotEmailEventInterface interface {
	Send()
}

type PasswordForgotEmailEvent struct {
	UserValidation entityInterface.UserValidationInterface
	Mail           mail.MailInterface
}

func NewPasswordForgotEmailEvent(userValidation entityInterface.UserValidationInterface) *PasswordForgotEmailEvent {
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
