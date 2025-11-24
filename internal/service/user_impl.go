package service

import (
	"go-crud/internal/dto/request"
	"go-crud/internal/dto/response"
	"go-crud/internal/repository"
)

type UserServiceImpl struct {
	Repository repository.UserRepository
}

func (s *UserServiceImpl) Create(user request.User) (response.User, error) {
	createdUser, err := s.Repository.Create(request.ToUserModel(user))
	return response.FromUserModel(createdUser), err
}

func (s *UserServiceImpl) GetAll() ([]response.User, error) {
	users, err := s.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	responseUsers := make([]response.User, len(users))
	for i, v := range users {
		responseUsers[i] = response.FromUserModel(v)
	}

	return responseUsers, nil
}

func (s *UserServiceImpl) GetById(id string) (response.User, error) {
	user, err := s.Repository.GetById(id)
	return response.FromUserModel(user), err
}

func (s *UserServiceImpl) Update(user request.User, id string) (response.User, error) {
	updatedUser, err := s.Repository.Update(request.ToUserModel(user), id)
	if err != nil {
		return response.User{}, err
	}

	return response.FromUserModel(updatedUser), nil
}

func (s *UserServiceImpl) Delete(id string) error {
	return s.Repository.Delete(id)
}
