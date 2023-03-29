package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/hpaes/api-project-golang/internal/dto"
	"github.com/hpaes/api-project-golang/internal/entity"
	"github.com/hpaes/api-project-golang/internal/infra/database"
)

type UserHandler struct {
	userDb       database.UserInterface
	jwt          *jwtauth.JWTAuth
	jwtExpiresIn int
}

func NewUserHandler(userDb database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		userDb:       userDb,
		jwt:          jwt,
		jwtExpiresIn: jwtExpiresIn,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// USECASE
	u, err := entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.userDb.Create(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetJwtInput(w http.ResponseWriter, r *http.Request) {
	var getJwtInput dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&getJwtInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// USECASE
	u, err := h.userDb.FindByEmail(getJwtInput.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if !u.ValidatePasswordHash(getJwtInput.Password) {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	token, err := generateToken(u, h)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func generateToken(u *entity.User, h *UserHandler) (string, error) {
	_, tokenString, err := h.jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.jwtExpiresIn)).Unix(),
	})

	if err != nil {
		return "", errors.New("error generating token")
	}

	return tokenString, nil
}
