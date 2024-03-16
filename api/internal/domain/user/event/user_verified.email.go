package event

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

type UserVerifiedEmailEventInterface interface {
	Send()
}

type UserVerifiedEmailEventInputDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserVerifiedEmailEvent struct {
	Data UserVerifiedEmailEventInputDTO
	Mail mail.MailInterface
}

func NewUserVerifiedEmailEvent(data UserVerifiedEmailEventInputDTO) *UserVerifiedEmailEvent {
	return &UserVerifiedEmailEvent{
		Data: data,
		Mail: mail.NewMail(),
	}
}

func (e *UserVerifiedEmailEvent) Send() {
	e.Mail.SetTo([]string{e.Data.Email})
	e.Mail.SetSubject("Seja muito bem vindo!")
	e.Mail.SetBody("user_verified", e.Data)
	err := e.Mail.Send()
	if err != nil {
		slog.Error("error sending email", err)
	}
}
