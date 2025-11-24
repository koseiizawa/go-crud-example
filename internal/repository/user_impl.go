package repository

import (
	"go-crud/internal/database"
	"go-crud/internal/model"
)

type UserRepositoryImpl struct{}

func (r *UserRepositoryImpl) Create(user model.User) (model.User, error) {
	err := database.DB.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetAll() ([]model.User, error) {
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetById(id string) (model.User, error) {
	var user model.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return model.User{}, nil
	}

	return user, nil
}

func (r *UserRepositoryImpl) Update(user model.User, id string) (model.User, error) {
	var existingUser model.User
	err := database.DB.First(&existingUser, id).Error
	if err != nil {
		return model.User{}, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email

	if err := database.DB.Save(&existingUser).Error; err != nil {
		return model.User{}, err
	}

	return existingUser, nil
}

func (r *UserRepositoryImpl) Delete(id string) error {
	if err := database.DB.Delete(model.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
