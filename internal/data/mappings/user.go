package mappings

import (
	"errors"

	"github.com/joaovds/chat/internal/data/dtos"
	"github.com/joaovds/chat/internal/domain/user"
)

const (
  ErrUserIsNil = "user is nil"
  ErrCreateUserInputDTOIsNil = "create user input dto is nil"
)

func MapCreateUserInputDTOToUser(createUserInputDTO *dtos.CreateUserInputDTO) (*user_domain.User, error) {
  if createUserInputDTO == nil {
    return nil, errors.New(ErrCreateUserInputDTOIsNil)
  }

  return &user_domain.User{
    Id:       createUserInputDTO.ID,
    Nickname: createUserInputDTO.Nickname,
    Password: createUserInputDTO.Password,
  }, nil
}

func MapUserToCreateUserInputDTO(user *user_domain.User) (*dtos.CreateUserInputDTO, error) {
  if user == nil {
    return nil, errors.New(ErrUserIsNil)
  }

  return &dtos.CreateUserInputDTO{
    ID:       user.Id,
    Nickname: user.Nickname,
    Password: user.Password,
  }, nil
}
