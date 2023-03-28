package entity

import (
	"apis/pkg/entity"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: hash,
	}, nil
}

func generatePassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash), nil
}

func (u *User) ValidatePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false, errors.New("invalid credentials")
	}

	return true, nil
}
