package services

import (
	"errors"
	"go-project/pkg/models"
	"go-project/pkg/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (us *UserService) CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	return us.UserRepo.CreateUser(user)
}

func (us *UserService) GetUserByID(id string) (*models.User, error) {
	return us.UserRepo.GetUserByID(id)
}

func (us *UserService) UpdateUser(user *models.User) error {
	if user.ID == "" {
		return errors.New("user id is required")
	}
	return us.UserRepo.UpdateUser(user)
}

func (us *UserService) DeleteUser(id string) error {
	return us.UserRepo.DeleteUser(id)
}

func (us *UserService) GetAllUsers() ([]*models.User, error) {
	return us.UserRepo.GetAllUsers()
}
