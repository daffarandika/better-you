package service

import (
	"betteryou/model"
	"betteryou/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService (repository *repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s UserService) GetAll() ([]model.User, error) {
	return s.repository.GetAll()
}

func (s UserService) GetByID(userID int) (*model.User, error) {
	return s.repository.GetByID(userID)
}

func (s UserService) Create(name string) error {
	newUser := &model.User {
		Name:	name,
	}
	return s.repository.Create(newUser)
}

func (s UserService) UpdateName(userID int, newName string) error {
	user, err := s.repository.GetByID(userID)
	if err != nil  {
		return err
	}
	user.Name = newName
	return s.repository.Update(userID, user)
}

func (s UserService) UpdateCoins(userID int, newCoinAmount int) error {
	user, err := s.repository.GetByID(userID)
	if err != nil  {
		return err
	}
	user.Coins = newCoinAmount
	return s.repository.Update(userID, user)
}

func (s UserService) AddCoins(userID int, amountAdded int) (int, error) {
	user, err := s.repository.GetByID(userID)
	if err != nil  {
		return -1, err
	}
	newAmount := user.Coins + amountAdded
	user.Coins = newAmount
	return newAmount, s.repository.Update(userID, user)
}

func (s UserService) DeleteByID(userID int) error {
	return s.repository.Delete(userID)
}
