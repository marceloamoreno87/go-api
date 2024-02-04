package entity

import (
	"errors"
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name string, email string, password string) (user *User, err error) {
	user = &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	valid := user.Validate()
	if valid != nil {
		return nil, valid
	}
	_, valid = user.IsEmailValid()
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
	return
}

func (u *User) IsEmailValid() (bool, error) {
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return false, errors.New("Email is invalid")
	}
	return true, err
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GetID() int64 {
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

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *User) SetID(id int64) {
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

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.CreatedAt = createdAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.UpdatedAt = updatedAt
}
