package inmemoryrepo

import (
	"github.com/google/uuid"
	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

type InMemoryUserRepo struct {
	users map[string]*models.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]*models.User),
	}
}
func (r *InMemoryUserRepo) CreateUser(input repository.CreateUserInput) error {
	newId := uuid.New().String()
	r.users[newId] = &models.User{
		ID:        newId,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Cart:      &models.Cart{},
	}
	return nil
}

func (r *InMemoryUserRepo) GetUser(input repository.GetUserInput) (*models.User, error) {
	user, ok := r.users[input.ID]
	if !ok {
		return nil, repository.ErrUserNotFound
	}
	return user, nil
}

func (r *InMemoryUserRepo) DeleteUser(input repository.DeleteUserInput) error {
	delete(r.users, input.ID)
	return nil
}

func (r *InMemoryUserRepo) ListUsers() ([]*models.User, error) {
	users := make([]*models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
