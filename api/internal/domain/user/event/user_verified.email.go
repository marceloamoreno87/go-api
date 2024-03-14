package event

import (
	"log/slog"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

type UserVerifiedEmailEventInterface interface {
	Send()
}

type UserVerifiedEmailEvent struct {
	UserValidation entityInterface.UserValidationInterface
	Mail           mail.MailInterface
}

func NewUserVerifiedEmailEvent(userValidation entityInterface.UserValidationInterface) *UserVerifiedEmailEvent {
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
