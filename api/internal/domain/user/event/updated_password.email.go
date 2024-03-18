package event

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

type UpdatedPasswordEmailEventInterface interface {
	Send()
}

type UpdatedPasswordEmailEventInputDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type UpdatedPasswordEmailEvent struct {
	Data UpdatedPasswordEmailEventInputDTO
	Mail mail.MailInterface
}

func NewUpdatedPasswordEmailEvent(data UpdatedPasswordEmailEventInputDTO) *UpdatedPasswordEmailEvent {
	return &UpdatedPasswordEmailEvent{
		Data: data,
		Mail: mail.NewMail(),
	}
}

func (e *UpdatedPasswordEmailEvent) Send() {
	e.Mail.SetTo([]string{e.Data.Email})
	e.Mail.SetSubject("Senha alterada")
	e.Mail.SetBody("internal/domain/user/views/forgot_password", e.Data)
	err := e.Mail.Send()
	if err != nil {
		slog.Error("error sending email", err)
	}
}
