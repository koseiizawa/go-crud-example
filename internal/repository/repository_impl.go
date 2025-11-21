package repository

import (
	"go-crud/internal/database"
	"go-crud/internal/domain"
	"go-crud/internal/model"
)

type UserRepositoryImpl struct{}

func (r *UserRepositoryImpl) Create(user domain.User) (domain.User, error) {
	createUser := UserDomainToModel(user)

	err := database.DB.Create(&createUser).Error
	if err != nil {
		return domain.User{}, err
	}

	return UserModelToDomain(createUser), nil
}

func (r *UserRepositoryImpl) GetAll() ([]domain.User, error) {
	var models []model.UserModel
	if err := database.DB.Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]domain.User, len(models))
	for i, v := range models {
		users[i] = UserModelToDomain(v)
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetById(id string) (domain.User, error) {
	var model model.UserModel
	err := database.DB.First(&model, id).Error
	if err != nil {
		return domain.User{}, nil
	}

	return UserModelToDomain(model), nil
}

func (r *UserRepositoryImpl) Update(user domain.User, id string) (domain.User, error) {
	var model model.UserModel
	err := database.DB.First(&model, id).Error
	if err != nil {
		return domain.User{}, err
	}

	model.Name = user.Name
	model.Email = user.Email

	if err := database.DB.Save(&model).Error; err != nil {
		return domain.User{}, err
	}

	return UserModelToDomain(model), nil
}

func (r *UserRepositoryImpl) Delete(id string) error {
	if err := database.DB.Delete(model.UserModel{}, id).Error; err != nil {
		return err
	}

	return nil
}

func UserDomainToModel(input domain.User) model.UserModel {
	return model.UserModel{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
}

func UserModelToDomain(input model.UserModel) domain.User {
	return domain.User{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}
}
