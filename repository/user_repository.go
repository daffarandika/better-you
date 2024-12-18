package repository

import (
	"betteryou/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository (db *gorm.DB) *UserRepository {
	return &UserRepository{
		db:	db,
	}
}

func (r UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r UserRepository) GetByID(userID int) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, userID)
	return &user, result.Error
}

func (r UserRepository) Create(user *model.User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r UserRepository) Update(userID int, updatedUser *model.User) error {
	result := r.db.Model(&model.User{}).Where("id = ?", userID).Updates(updatedUser)
	return result.Error
}

func (r UserRepository) Delete(userID int) error {
	user, err := r.GetByID(userID)
	if err != nil {
		return err
	}
	result := r.db.Delete(user)
	return result.Error
}
