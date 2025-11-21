package service

import (
	"go-crud/internal/dto"
)

type UserService interface {
	Create(user dto.UserRequest) (dto.UserResponse, error)
	GetAll() ([]dto.UserResponse, error)
	GetById(id string) (dto.UserResponse, error)
	Update(user dto.UserRequest, id string) (dto.UserResponse, error)
	Delete(id string) error
}
