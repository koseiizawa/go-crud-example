package repository

import (
	"errors"
	"go-crud/internal/database"
	"go-crud/internal/model"
)

type AdminRepositoryImpl struct{}

func (r *AdminRepositoryImpl) SignUp(admin model.Admin) (model.Admin, error) {
	err := database.DB.Create(&admin).Error
	if err != nil {
		return model.Admin{}, err
	}

	return admin, nil
}

func (r *AdminRepositoryImpl) LogIn(admin model.Admin) (model.Admin, error) {
	email := admin.Email
	var existedUser model.Admin
	err := database.DB.Where("email = ?", email).First(&existedUser).Error
	if err != nil {
		return model.Admin{}, err
	}

	if admin.Password != existedUser.Password {
		return model.Admin{}, errors.New("authentication failed")
	}

	return existedUser, nil
}
