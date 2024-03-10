package event

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

type UpdatedPasswordEmailEventInterface interface {
	Send()
}

type UpdatedPasswordEmailEvent struct {
	user entity.UserInterface
	Mail mail.MailInterface
}

func NewUpdatedPasswordEmailEvent(user entity.UserInterface) *UpdatedPasswordEmailEvent {
	return &UpdatedPasswordEmailEvent{
		user: user,
		Mail: mail.NewMail(),
	}
}

func (e *UpdatedPasswordEmailEvent) Send() {
	e.Mail.SetTo([]string{e.user.GetEmail()})
	e.Mail.SetSubject("Senha alterada")
	e.Mail.SetBody("forgot_password", e.user)
	err := e.Mail.Send()
	if err != nil {
		slog.Error("error sending email", err)
	}
}
