package service

import (
	"go-crud/internal/dto/request"
	"go-crud/internal/dto/response"
)

type AdminService interface {
	SignUp(admin request.Admin) (response.Admin, error)
	LogIn(admin request.Admin) (response.Admin, error)
}
