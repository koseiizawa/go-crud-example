package repository

import "go-crud/internal/domain"

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	GetAll() ([]domain.User, error)
	GetById(id string) (domain.User, error)
	Update(user domain.User, id string) (domain.User, error)
	Delete(id string) error
}
