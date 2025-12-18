package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/luis13005/pos-go/internal/dto"
	"github.com/luis13005/pos-go/internal/infra/database"
	"github.com/luis13005/pos-go/internal/model"
)

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDB:       db,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserDto
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

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWT
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	usuarioEncontrado, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isValid, err := usuarioEncontrado.ValidaSenha(user.Senha)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if isValid == false {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("senha inválida"))
		return
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": usuarioEncontrado.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})

	accessToken := struct {
		AccesToken string `json"access_token"`
	}{
		AccesToken: tokenString,
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandler) FindUserByEmail(w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(u.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
