package service

import (
	"go-crud/internal/dto/request"
	"go-crud/internal/dto/response"
)

type UserService interface {
	Create(user request.User) (response.User, error)
	GetAll() ([]response.User, error)
	GetById(id string) (response.User, error)
	Update(user request.User, id string) (response.User, error)
	Delete(id string) error
}
