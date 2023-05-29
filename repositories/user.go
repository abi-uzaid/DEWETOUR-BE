package repositories

import (
	"dewetour/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Find(&user, ID).Error

	return user, err
}

func (r *userRepository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
