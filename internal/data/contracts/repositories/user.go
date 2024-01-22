package repositories_data

import "github.com/joaovds/chat/internal/data/dtos"

type CreateUserRepository interface {
  Create(user *dtos.CreateUserInputDTO) error
}
