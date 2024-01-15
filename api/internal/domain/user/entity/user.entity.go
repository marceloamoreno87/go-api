package entity

import "errors"

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Photo     string `json:"photo"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUser(username string, password string, photo string) (user *User, err error) {
	user = &User{
		Username: username,
		Password: password,
		Photo:    photo,
	}
	valid := user.Validate()
	if valid != nil {
		return
	}
	return
}

func (u *User) Validate() (err error) {
	if u.Username == "" {
		return errors.New("Username is required")
	}
	if u.Password == "" {
		return errors.New("Password is required")
	}
	return
}

func (u *User) GetUserName() string {
	return u.Username
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetPhoto() string {
	return u.Photo
}

func (u *User) GetCreatedAt() string {
	return u.CreatedAt
}

func (u *User) GetUpdatedAt() string {
	return u.UpdatedAt
}

func (u *User) SetUserName(username string) {
	u.Username = username
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetPhoto(photo string) {
	u.Photo = photo
}

func (u *User) SetCreatedAt(createdAt string) {
	u.CreatedAt = createdAt
}

func (u *User) SetUpdatedAt(updatedAt string) {
	u.UpdatedAt = updatedAt
}
