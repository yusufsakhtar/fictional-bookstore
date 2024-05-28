package inmemoryrepo

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

type InMemoryUserRepo struct {
	users map[string]*models.User
	mu    sync.RWMutex
}

// NewInMemoryUserRepo creates a new InMemoryUserRepo, optionally seeding it from a file
func NewInMemoryUserRepo(seedFromFile bool, seedFileName string) *InMemoryUserRepo {
	if seedFromFile {
		file, err := os.Open(seedFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		users := make(map[string]*models.User)
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&users)
		if err != nil {
			log.Fatal(err)
		}

		return &InMemoryUserRepo{
			users: users,
		}
	} else {
		return &InMemoryUserRepo{
			users: make(map[string]*models.User),
		}
	}
}

func (r *InMemoryUserRepo) CreateUser(input repository.CreateUserInput) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	newId := uuid.New().String()
	r.users[newId] = &models.User{
		ID:        newId,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	return r.users[newId], nil
}

func (r *InMemoryUserRepo) GetUser(input repository.GetUserInput) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[input.ID]
	if !ok {
		return nil, repository.ErrUserNotFound
	}
	return user, nil
}

func (r *InMemoryUserRepo) DeleteUser(input repository.DeleteUserInput) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	delete(r.users, input.ID)
	return nil
}

func (r *InMemoryUserRepo) ListUsers() ([]*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*models.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
