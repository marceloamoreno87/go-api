package entity

import (
	"errors"
	"net/mail"
	"time"

	"github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int32        `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	RoleID    int32        `json:"role_id"`
	AvatarID  int32        `json:"avatar_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Role      *entity.Role `json:"role"`
	Avatar    *Avatar      `json:"avatar"`
}

func NewUser(name string, email string, password string, roleID int32, avatarID int32) (user *User, err error) {
	user = &User{
		Name:     name,
		Email:    email,
		Password: password,
		RoleID:   roleID,
		AvatarID: avatarID,
	}
	valid := user.Validate()
	if valid != nil {
		return nil, valid
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.SetPassword(string(hash))

	return
}

func (u *User) Validate() (err error) {
	if u.Name == "" {
		return errors.New("Name is required")
	}
	if u.Email == "" {
		return errors.New("Email is required")
	}
	if u.Password == "" {
		return errors.New("Password is required")
	}
	if u.RoleID == 0 {
		return errors.New("Role is required")
	}
	if u.AvatarID == 0 {
		return errors.New("Avatar is required")
	}
	_, err = mail.ParseAddress(u.Email)
	if err != nil {
		return errors.New("Email is invalid")
	}
	return
}

func (u *User) ComparePassword(password string) (b bool) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
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

func (u *User) GetRole() *entity.Role {
	return u.Role
}

func (u *User) GetAvatar() *Avatar {
	return u.Avatar
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

func (u *User) SetRole(role *entity.Role) {
	u.Role = role
}

func (u *User) SetAvatar(avatar *Avatar) {
	u.Avatar = avatar
}
