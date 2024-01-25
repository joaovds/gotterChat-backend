package webserver_mappings

import (
	"errors"

	user_domain "github.com/joaovds/chat/internal/domain/user"
	webserver_dtos "github.com/joaovds/chat/internal/infra/webserver/dtos"
)

const (
	ErrInvalidUser = "Invalid user"
)

func MapCreateUserRequestDTOToUser(createUserRequestDTO *webserver_dtos.CreateUserRequestDTO) (*user_domain.User, error) {
	if createUserRequestDTO == nil {
		return nil, errors.New(ErrInvalidUser)
	}

	return &user_domain.User{
		Nickname: createUserRequestDTO.Nickname,
		Password: createUserRequestDTO.Password,
	}, nil
}
