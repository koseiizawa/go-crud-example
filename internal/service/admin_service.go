package service

import "go-crud/internal/dto"

type AdminService interface {
	SignUp(user dto.AdminRequest) (dto.AdminResponse, error)
	LogIn(user dto.AdminRequest) (dto.AdminResponse, error)
}
