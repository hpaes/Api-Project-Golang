package entity

import (
	"errors"
	"log"
	"strings"
	"unicode"

	"github.com/hpaes/api-project-golang/pkg/entity"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	err := validateNameAndEmail(name, email)
	if err != nil {
		return nil, err
	}
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

func (u *User) ValidatePasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func validateNameAndEmail(name, email string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	if email == "" {
		return errors.New("email cannot be empty")
	}

	return nil
}

func generatePassword(password string) (string, error) {
	password, err := validatePassword(password)
	if err != nil {
		return "", err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash), nil
}

func validatePassword(password string) (string, error) {
	if len(password) < 8 {
		return "", errors.New("password must be at least 8 characters")
	}

	password = strings.TrimSpace(password)

	letters := 0
	number := false
	upper := false
	special := false
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			letters++
		case unicode.IsLower(c):
			letters++
		case unicode.IsLetter(c):
			letters++
		case unicode.IsSpace(c):
			return "", errors.New("password must not contain spaces")
		default:
			return "", errors.New("password must contain at least one number, one uppercase letter, one lowercase letter and one special character")

		}
	}

	isValid := number && upper && special && letters >= 8

	if isValid {
		return password, nil
	}

	return "", errors.New("password must contain at least one number, one uppercase letter, one lowercase letter and one special character")
}
