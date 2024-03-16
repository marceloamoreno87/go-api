package entity

import (
	"errors"
	"net/mail"
	"time"

	entityInterface "github.com/marceloamoreno/goapi/internal/domain/user/interface/entity"
	"github.com/marceloamoreno/goapi/internal/shared/notification"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Active    bool      `json:"active"`
	RoleID    int32     `json:"role_id"`
	AvatarID  int32     `json:"avatar_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name string, email string, password string) (user entityInterface.UserInterface, err error) {
	user = &User{
		Name:     name,
		Email:    email,
		Password: password,
		RoleID:   2,
		AvatarID: 1,
		Active:   false,
	}

	notify := user.Validate()
	if notify.HasErrors() {
		return nil, errors.New(notify.Messages())
	}

	return
}

func (u *User) Validate() (notify notification.ErrorsInterface) {

	notify = notification.New()
	if u.Name == "" {
		notify.AddError("Name is required", "user.entity.name")
	}
	if u.Email == "" {
		notify.AddError("Email is required", "user.entity.email")
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		notify.AddError("Email is invalid", "user.entity.email")
	}
	if u.Password == "" {
		notify.AddError("Password is required", "user.entity.password")
	}
	if u.RoleID == 0 {
		notify.AddError("Role is required", "user.entity.role_id")
	}
	if u.AvatarID == 0 {
		notify.AddError("Avatar is required", "user.entity.avatar_id")
	}
	return
}

func (u *User) ComparePassword(password string) (b bool) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) HashPassword() {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u.Password = string(hashed)
}

func (u *User) GetID() int32 {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetActive() bool {
	return u.Active
}

func (u *User) GetRoleID() int32 {
	return u.RoleID
}

func (u *User) GetAvatarID() int32 {
	return u.AvatarID
}

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *User) SetID(id int32) {
	u.ID = id
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetActive(active bool) {
	u.Active = active
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetRoleID(roleID int32) {
	u.RoleID = roleID
}

func (u *User) SetAvatarID(avatarID int32) {
	u.AvatarID = avatarID
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}
