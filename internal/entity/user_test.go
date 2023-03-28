package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenValidParamShouldCreateUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@email.com", "password")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@email.com", user.Email)

	isValid, err := user.ValidatePassword("password")
	assert.Equal(t, true, isValid)
	assert.Nil(t, err)
	assert.NotEqual(t, "password", user.Password)
}

func TestGivenInvalidPasswordShouldReturnError(t *testing.T) {
	user, err := NewUser("John Doe", "john@email.com", "password")
	assert.Nil(t, err)
	isValid, err := user.ValidatePassword("")
	assert.Error(t, err)
	assert.Equal(t, false, isValid)
}

func TestGivenInvalidParamsShouldNotCreateUser(t *testing.T) {
	user, err := NewUser("", "", "")
	assert.Error(t, err, "password cannot be empty")
	assert.Nil(t, user)
}
