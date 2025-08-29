package repositories

import (
	"errors"
	"go-project/pkg/models"
	"sync"
)

type UserRepository struct {
	mu    sync.Mutex
	users map[string]*models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[string]*models.User)}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.users[user.ID]
	if !exists {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) DeleteUser(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, exists := r.users[id]
	if !exists {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}

func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	users := make([]*models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
