package event

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

type UserVerifyEmailEventInterface interface {
	Send()
}

type UserVerifyEmailEventInputDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Hash  string `json:"hash"`
}

type UserVerifyEmailEvent struct {
	Data UserVerifyEmailEventInputDTO
	Mail mail.MailInterface
}

func NewUserVerifyEmailEvent(data UserVerifyEmailEventInputDTO) *UserVerifyEmailEvent {
	return &UserVerifyEmailEvent{
		Data: data,
		Mail: mail.NewMail(),
	}
}

func (e *UserVerifyEmailEvent) Send() {
	e.Mail.SetTo([]string{e.Data.Email})
	e.Mail.SetSubject("Ative sua conta!")
	e.Mail.SetBody("user_verify", e.Data)
	err := e.Mail.Send()
	if err != nil {
		slog.Error("error sending email", err)
	}
}
