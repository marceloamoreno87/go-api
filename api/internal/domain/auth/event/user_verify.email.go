package event

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/auth/entity"
	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

type UserVerifyEmailEventInterface interface {
	Send()
}

type UserVerifyEmailEvent struct {
	UserValidation entity.UserValidationInterface
	Mail           mail.MailInterface
}

func NewUserVerifyEmailEvent(userValidation entity.UserValidationInterface) *UserVerifyEmailEvent {
	return &UserVerifyEmailEvent{
		UserValidation: userValidation,
		Mail:           mail.NewMail(),
	}
}

func (e *UserVerifyEmailEvent) Send() {
	e.Mail.SetTo([]string{e.UserValidation.GetUser().GetEmail()})
	e.Mail.SetSubject("Ative sua conta!")
	e.Mail.SetBody("user_verify", e.UserValidation)
	err := e.Mail.Send()
	if err != nil {
		slog.Error("error sending email", err)
	}
}
