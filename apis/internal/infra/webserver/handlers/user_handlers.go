package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gabrielsgradinar/golang-learning/apis/internal/dto"
	"github.com/gabrielsgradinar/golang-learning/apis/internal/entity"
	"github.com/gabrielsgradinar/golang-learning/apis/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct{
	UserDB database.UserInterface
	Jwt *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, JwtExpiresIn int) *UserHandler {
	return &UserHandler{ 
		UserDB: db,
		Jwt: jwt,
		JwtExpiresIn: JwtExpiresIn,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request){
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !u.ValidadePassword(user.Password){
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]interface{} {
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})
	accesToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accesToken)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request){
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}