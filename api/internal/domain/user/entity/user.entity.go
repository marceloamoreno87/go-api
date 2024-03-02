package entity

import (
	"errors"
	"net/mail"
	"time"

	AvatarEntity "github.com/marceloamoreno/goapi/internal/domain/avatar/entity"
	RoleEntity "github.com/marceloamoreno/goapi/internal/domain/role/entity"
	"github.com/marceloamoreno/goapi/internal/shared/notification"
	"github.com/marceloamoreno/goapi/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int32                `json:"id"`
	Name      string               `json:"name"`
	Email     string               `json:"email"`
	Password  string               `json:"password"`
	Token     string               `json:"token"`
	RoleID    int32                `json:"role_id"`
	AvatarID  int32                `json:"avatar_id"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
	Role      *RoleEntity.Role     `json:"role"`
	Avatar    *AvatarEntity.Avatar `json:"avatar"`
}

func NewUser(name string, email string, password string, roleID int32, avatarID int32) (user *User, err error) {
	user = &User{
		Name:     name,
		Email:    email,
		Password: password,
		RoleID:   roleID,
		AvatarID: avatarID,
	}
	notify := user.Validate()
	if notify.HasErrors() {
		return nil, errors.New(notify.Messages())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.SetPassword(string(hash))
	return
}

func (u *User) Validate() (notify *notification.Errors) {

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

func (u *User) GenerateToken() (err error) {
	claims := map[string]interface{}{
		"id":        u.ID,
		"name":      u.Name,
		"email":     u.Email,
		"role_id":   u.RoleID,
		"avatar_id": u.AvatarID,
	}
	jwtauth := jwt.New(claims)
	u.Token = jwtauth.Token
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

func (u *User) GetRole() *RoleEntity.Role {
	return u.Role
}

func (u *User) GetAvatar() *AvatarEntity.Avatar {
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

func (u *User) SetRole(role *RoleEntity.Role) {
	u.Role = role
}

func (u *User) SetAvatar(avatar *AvatarEntity.Avatar) {
	u.Avatar = avatar
}
