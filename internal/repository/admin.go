package repository

import (
	"go-crud/internal/model"
)

type AdminRepository interface {
	SignUp(admin model.Admin) (model.Admin, error)
	LogIn(admin model.Admin) (model.Admin, error)
}
