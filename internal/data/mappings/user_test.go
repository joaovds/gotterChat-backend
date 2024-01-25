package mappings

import (
	"testing"

	"github.com/joaovds/chat/internal/data/dtos"
	user_domain "github.com/joaovds/chat/internal/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestMapCreateUserInputDTOToUser(t *testing.T) {
	t.Run("should map create user input dto to user", func(t *testing.T) {
		createUserInputDTO := &dtos.CreateUserInputDTO{
			ID:       "123",
			Nickname: "nickname",
			Password: "password",
		}

		user, err := MapCreateUserInputDTOToUser(createUserInputDTO)

		assert.Equal(t, createUserInputDTO.ID, user.Id)
		assert.Equal(t, createUserInputDTO.Nickname, user.Nickname)
		assert.Equal(t, createUserInputDTO.Password, user.Password)
		assert.IsType(t, &user_domain.User{}, user)
		assert.Nil(t, err)
	})

	t.Run("should return error if create user input dto is nil", func(t *testing.T) {
		user, err := MapCreateUserInputDTOToUser(nil)
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, "create user input dto is nil", err.Error())
	})
}

func TestMapUserToCreateUserInputDTO(t *testing.T) {
	t.Run("should map user to create user input dto", func(t *testing.T) {
		user := &user_domain.User{
			Id:       "123",
			Nickname: "nickname",
			Password: "password",
		}

		createUserInputDTO, err := MapUserToCreateUserInputDTO(user)

		assert.Equal(t, user.Id, createUserInputDTO.ID)
		assert.Equal(t, user.Nickname, createUserInputDTO.Nickname)
		assert.Equal(t, user.Password, createUserInputDTO.Password)
		assert.IsType(t, &dtos.CreateUserInputDTO{}, createUserInputDTO)
		assert.Nil(t, err)
	})

	t.Run("should return error if user is nil", func(t *testing.T) {
		createUserInputDTO, err := MapUserToCreateUserInputDTO(nil)

		assert.Nil(t, createUserInputDTO)
		assert.NotNil(t, err)
		assert.Equal(t, "user is nil", err.Error())
	})
}
