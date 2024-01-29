package user_data_usecase

import (
	"testing"

	"github.com/joaovds/chat/internal/data/dtos"
	"github.com/joaovds/chat/internal/data/mappings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCreateUserUseCase(t *testing.T) {
	t.Run("Should return a CreateUserUseCase", func(t *testing.T) {
		createUserUseCase := NewCreateUserUseCase(&CreateUserRepositoryMock{})

		assert.NotNil(t, createUserUseCase)
		assert.IsType(t, &CreateUserUseCase{}, createUserUseCase)
	})
}

func TestCreateUserUseCase_Execute(t *testing.T) {
	t.Run("Should return an error when mapping user to CreateUserInputDTO", func(t *testing.T) {
    createUserRepositoryMock := &CreateUserRepositoryMock{}
		createUserUseCase := NewCreateUserUseCase(createUserRepositoryMock)

		err := createUserUseCase.Execute(nil)

		assert.NotNil(t, err)
		assert.Equal(t, mappings.ErrUserIsNil, err.Error())
    createUserRepositoryMock.AssertNumberOfCalls(t, "Create", 0)
	})
}

type CreateUserRepositoryMock struct {
	mock.Mock
}

func (mock *CreateUserRepositoryMock) Create(user *dtos.CreateUserInputDTO) error {
	args := mock.Called(user)
	return args.Error(0)
}
