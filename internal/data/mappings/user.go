package mappings

import (
	"github.com/joaovds/chat/internal/domain/user"
	"github.com/joaovds/chat/internal/data/dtos"
)

func MapCreateUserInputDTOToUser(createUserInputDTO *dtos.CreateUserInputDTO) *user_domain.User {
  return &user_domain.User{
    Id:       createUserInputDTO.ID,
    Nickname: createUserInputDTO.Nickname,
    Password: createUserInputDTO.Password,
  }
}

func MapUserToCreateUserInputDTO(user *user_domain.User) *dtos.CreateUserInputDTO {
  return &dtos.CreateUserInputDTO{
    ID:       user.Id,
    Nickname: user.Nickname,
    Password: user.Password,
  }
}
