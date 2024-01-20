package user_data_usecase

import (
	repositories_data "github.com/joaovds/chat/internal/data/contracts/repositories"
	"github.com/joaovds/chat/internal/data/mappings"
	"github.com/joaovds/chat/internal/domain/user"
)

type CreateUserUseCase struct {
  createUserRepository repositories_data.CreateUserRepository
}

func NewCreateUserUseCase(createUserRepository repositories_data.CreateUserRepository) *CreateUserUseCase {
  return &CreateUserUseCase{createUserRepository: createUserRepository}
}

func (createUserUseCase *CreateUserUseCase) Execute(user *user_domain.User) error {
  mappedUser := mappings.MapUserToCreateUserInputDTO(user)

  return createUserUseCase.createUserRepository.Create(mappedUser)
}
