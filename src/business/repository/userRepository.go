package repository

import (
	"jwt-restApi/src/business/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.User) error
	FindUserByEmail(email string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user entity.User) error {
	err := r.db.Create(&user).Error
	return err
}

func (r *userRepository) FindUserByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "email = ?", email).Error
	return user, err
}
