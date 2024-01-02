package database

import "github.com/gabrielsgradinar/golang-learning/apis/internal/entity"


type UserInterface interface{
	Create(user entity.User)
	FindByEmail(email string)(*entity.User, error)
}