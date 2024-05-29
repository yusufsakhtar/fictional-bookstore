package service

import (
	"github.com/yusufsakhtar/playstation-assignment/internal/models"
	"github.com/yusufsakhtar/playstation-assignment/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepo
	cartRepo repository.CartRepo
}

func NewUserService(userRepo repository.UserRepo, cartRepo repository.CartRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
		cartRepo: cartRepo,
	}
}

func (s *UserService) ListUsers() ([]*models.User, error) {
	return s.userRepo.ListUsers()
}

func (s *UserService) CreateUser(input repository.CreateUserInput) error {
	user, err := s.userRepo.CreateUser(input)
	if err != nil {
		return err
	}
	err = s.cartRepo.CreateCart(repository.CreateCartInput{UserID: user.ID})
	return err
}

func (s *UserService) GetUser(input repository.GetUserInput) (*models.User, error) {
	return s.userRepo.GetUser(input)
}

// TODO: move this to  cart service (cart svc didnt exist initially, this code felt like an ok fit here for a time)
func (s *UserService) GetUserCart(input repository.GetUserCartInput) (*models.Cart, error) {
	cart, err := s.cartRepo.GetUserCart(input)
	if err != nil {
		// only doing this for bootstrapping purposes
		// normally we would return the error as is
		if err == repository.ErrCartNotFound {
			err = s.cartRepo.CreateCart(repository.CreateCartInput{UserID: input.UserID})
			if err != nil {
				return nil, err
			}
			return s.cartRepo.GetUserCart(input)
		}
		return nil, err
	}
	return cart, nil
}

func (s *UserService) DeleteUser(input repository.DeleteUserInput) error {
	return s.userRepo.DeleteUser(input)
}
