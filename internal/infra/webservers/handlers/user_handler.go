package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hpaes/api-project-golang/internal/dto"
	"github.com/hpaes/api-project-golang/internal/entity"
	"github.com/hpaes/api-project-golang/internal/infra/database"
)

type UserHandler struct {
	userDb database.UserInterface
}

func NewUserHandler(userDb database.UserInterface) *UserHandler {
	return &UserHandler{
		userDb: userDb,
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
