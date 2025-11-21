package repository

import (
	"errors"
	"go-crud/internal/database"
	"go-crud/internal/domain"
	"go-crud/internal/model"
)

type AdminRepositoryImpl struct{}

func (r *AdminRepositoryImpl) SignUp(user domain.AdminUser) (domain.AdminUser, error) {
	createdUser := AdminDomainToModel(user)

	err := database.DB.Create(&createdUser).Error
	if err != nil {
		return domain.AdminUser{}, err
	}

	return AdminModelToDomain(createdUser), nil
}

func (r *AdminRepositoryImpl) LogIn(user domain.AdminUser) (domain.AdminUser, error) {
	email := user.Email
	var existedUser model.Admin
	err := database.DB.Where("email = ?", email).First(&existedUser).Error
	if err != nil {
		return domain.AdminUser{}, err
	}

	if user.Password != existedUser.Password {
		return domain.AdminUser{}, errors.New("authentication failed")
	}

	return AdminModelToDomain(existedUser), nil
}

func AdminDomainToModel(input domain.AdminUser) model.Admin {
	return model.Admin{
		Email:    input.Email,
		Password: input.Password,
	}
}

func AdminModelToDomain(input model.Admin) domain.AdminUser {
	return domain.AdminUser{
		ID:    input.ID,
		Email: input.Email,
	}
}
