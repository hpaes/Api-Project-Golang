package database

import (
	"testing"

	"github.com/hpaes/api-project-golang/internal/entity"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGivenValidParamsShouldInsertUserInDb(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.User{})

	userRepo := NewUserRepository(db)
	user, err := entity.NewUser("John Doe", "john@email.com", "Passw0rd#")
	assert.NoError(t, err)

	err = userRepo.Create(user)
	assert.NoError(t, err)

	var userFound entity.User
	err = db.Find(&userFound, "id = ?", user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestGivenValidEmailShouldReturnUserFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.User{})

	userRepo := NewUserRepository(db)
	user, err := entity.NewUser("John Doe", "john@email.com", "passWord123#")
	assert.NoError(t, err)

	err = userRepo.Create(user)
	assert.NoError(t, err)

	userFound, err := userRepo.FindByEmail(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestGivenNonExistingEmailShouldNotReturnUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.User{})

	userRepo := NewUserRepository(db)
	user, err := entity.NewUser("John Doe", "john@email.com", "passWord123#")
	assert.NoError(t, err)

	err = userRepo.Create(user)
	assert.NoError(t, err)

	userFound, err := userRepo.FindByEmail("abc")
	assert.Error(t, err)
	assert.Nil(t, userFound)
}
