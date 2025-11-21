package repository

import "go-crud/internal/domain"

type AdminRepository interface {
	SignUp(user domain.AdminUser) (domain.AdminUser, error)
	LogIn(user domain.AdminUser) (domain.AdminUser, error)
}
