package mappings

import (
	"errors"

	"github.com/joaovds/chat/internal/data/dtos"
	"github.com/joaovds/chat/internal/domain/user"
)

func MapCreateUserInputDTOToUser(createUserInputDTO *dtos.CreateUserInputDTO) (*user_domain.User, error) {
  if createUserInputDTO == nil {
    return nil, errors.New("create user input dto is nil")
  }

  return &user_domain.User{
    Id:       createUserInputDTO.ID,
    Nickname: createUserInputDTO.Nickname,
    Password: createUserInputDTO.Password,
  }, nil
}

func MapUserToCreateUserInputDTO(user *user_domain.User) (*dtos.CreateUserInputDTO, error) {
  if user == nil {
    return nil, errors.New("user is nil")
  }

  return &dtos.CreateUserInputDTO{
    ID:       user.Id,
    Nickname: user.Nickname,
    Password: user.Password,
  }, nil
}
