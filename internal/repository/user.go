package repository

import (
	"go-crud/internal/model"
)

type UserRepository interface {
	Create(user model.User) (model.User, error)
	GetAll() ([]model.User, error)
	GetById(id string) (model.User, error)
	Update(user model.User, id string) (model.User, error)
	Delete(id string) error
}
