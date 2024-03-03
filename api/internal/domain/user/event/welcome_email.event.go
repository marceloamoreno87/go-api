package event

import (
	"log/slog"

	"github.com/marceloamoreno/goapi/internal/domain/user/entity"
	"github.com/marceloamoreno/goapi/internal/shared/mail"
)

type WelcomeEmailEvent struct {
	User entity.User `json:"user"`
}

func NewWelcomeEmailEvent(user entity.User) *WelcomeEmailEvent {
	return &WelcomeEmailEvent{
		User: user,
	}
}

func (e *WelcomeEmailEvent) Send() {
	m := mail.NewMail()
	m.SetTo([]string{e.User.Email})
	m.SetSubject("Seja muito bem vindo!")
	m.SetBody("welcome", e.User)
	err := m.Send()
	if err != nil {
		slog.Error("error sending email", err)
	}
}
