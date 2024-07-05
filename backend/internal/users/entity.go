package users

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"
)

var (
	BlankPassword = fmt.Sprintf("%v", md5.Sum([]byte("")))

	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordLength   = errors.New("password must be at least 8 characters")
	ErrLoginRequired    = errors.New("login is required")
	ErrNameRequired     = errors.New("name is required")
)

func New(name, login, password string) (*User, error) {
	now := time.Now()

	u := User{
		Name: name, Login: login, ModifiedAt: now,
	}

	if err := u.SetPassword(password); err != nil {
		return nil, err
	}

	return &u, nil
}

type User struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	DeletedAt   time.Time `json:"-"`
	LastLoginAt time.Time `json:"last_login_at"`
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	if len(password) < 8 {
		return ErrPasswordLength
	}

	u.Password = fmt.Sprintf("%x", (md5.Sum([]byte(password))))
	return nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return ErrNameRequired
	}

	if u.Login == "" {
		return ErrLoginRequired
	}

	if u.Password == BlankPassword {
		return ErrPasswordRequired
	}

	return nil
}
