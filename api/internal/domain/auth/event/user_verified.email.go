package event

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/auth/entity"
	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

type UserVerifiedEmailEventInterface interface {
	Send()
}

type UserVerifiedEmailEvent struct {
	UserValidation entity.UserValidationInterface
	Mail           mail.MailInterface
}

func NewUserVerifiedEmailEvent(userValidation entity.UserValidationInterface) *UserVerifiedEmailEvent {
	return &UserVerifiedEmailEvent{
		UserValidation: userValidation,
		Mail:           mail.NewMail(),
	}
}

func (e *UserVerifiedEmailEvent) Send() {
	e.Mail.SetTo([]string{e.UserValidation.GetUser().GetEmail()})
	e.Mail.SetSubject("Seja muito bem vindo!")
	e.Mail.SetBody("user_verified", e.UserValidation)
	err := e.Mail.Send()
	if err != nil {
		slog.Error("error sending email", err)
	}
}
