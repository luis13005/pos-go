package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luis13005/pos-go/internal/dto"
	"github.com/luis13005/pos-go/internal/infra/database"
	"github.com/luis13005/pos-go/internal/model"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: db}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.UserDto
	json.NewDecoder(r.Body).Decode(&user)

	newUser, err := model.NewUser(&model.User{Nome: user.Nome, Email: user.Email, Senha: user.Senha})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = h.UserDB.CreateUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
