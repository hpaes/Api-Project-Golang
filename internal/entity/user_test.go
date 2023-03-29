package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenValidParamShouldCreateUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@email.com", "passWord123$")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@email.com", user.Email)

	isValid := user.ValidatePasswordHash("passWord123$")
	assert.Equal(t, true, isValid)
	assert.NotEqual(t, "passWord123$", user.Password)
}

func TestGivenInvalidPasswordCredentialShouldReturnError(t *testing.T) {
	user, err := NewUser("John Doe", "john@email.com", "passWord123$")
	assert.Nil(t, err)
	isValid := user.ValidatePasswordHash("")
	assert.Equal(t, false, isValid)
}

func TestGivenInvalidNameShouldNotCreateUser(t *testing.T) {
	user, err := NewUser("", "john@email.com", "passWord123$")
	assert.EqualError(t, err, "name cannot be empty")
	assert.Nil(t, user)
}

func TestGivenInvalidEmailShouldNotCreateUser(t *testing.T) {
	user, err := NewUser("John Doe", "", "passWord123$")
	assert.EqualError(t, err, "email cannot be empty")
	assert.Nil(t, user)
}

func TestGivenShortOrEmptyPasswordShouldReturnError(t *testing.T) {
	user, err := NewUser("John Doe", "john@email.com", "")
	assert.EqualError(t, err, "password must be at least 8 characters")
	assert.Nil(t, user)
}

func TestGivenInvalidPasswordShouldReturnError(t *testing.T) {
	user, err := NewUser("John Doe", "john@email.com", "password")
	assert.EqualError(t, err, "password must contain at least one number, one uppercase letter, one lowercase letter and one special character")
	assert.Nil(t, user)
}

func TestGivenPasswordWithSpacesShouldReturnError(t *testing.T) {
	user, err := NewUser("John Doe", "john@email.com", "pass Word123$")
	assert.EqualError(t, err, "password must not contain spaces")
	assert.Nil(t, user)
}
