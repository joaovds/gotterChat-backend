package user_domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
  t.Run("should return a user", func(t *testing.T) {
    user, err := NewUser(uuid.New().String(), "nickname", "password")

    assert.Nil(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, "nickname", user.Nickname)
    assert.Equal(t, "password", user.Password)
    assert.NotNil(t, user.Id)
  })

  t.Run("should return an error when attribute is empty", func(t *testing.T) {
    _, err := NewUser("", "nickname", "password")
    assert.NotNil(t, err)
    assert.Equal(t, "invalid id", err.Error())

    _, err = NewUser(uuid.New().String(), "", "password")
    assert.NotNil(t, err)
    assert.Equal(t, "invalid nickname", err.Error())

    _, err = NewUser(uuid.New().String(), "nickname", "")
    assert.NotNil(t, err)
    assert.Equal(t, "invalid password", err.Error())
  })

  t.Run("should return an error when id is not a valid uuid", func(t *testing.T) {
    _, err := NewUser("invalid-uuid", "nickname", "password")

    assert.NotNil(t, err)
    assert.Equal(t, "id is not a valid uuid", err.Error())
  })
}
