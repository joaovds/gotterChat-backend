package user_domain

import (
	"errors"

	"github.com/joaovds/chat/pkg/validation"
)

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func NewUser(id, nickname, password string) (*User, error) {
	user := &User{
		Id:       id,
		Nickname: nickname,
		Password: password,
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) CreateInvalidError(attribute string) error {
	return errors.New("invalid " + attribute)
}

func (u *User) Validate() error {
	if u.Id == "" {
		return u.CreateInvalidError("id")
	}

	if !validation.IsValidUUID(u.Id) {
		return validation.CreateInvalidError("id")
	}

	if u.Nickname == "" {
		return u.CreateInvalidError("nickname")
	}

	if u.Password == "" {
		return u.CreateInvalidError("password")
	}

	return nil
}
