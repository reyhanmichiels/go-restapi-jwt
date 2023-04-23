package repository

import "gorm.io/gorm"

type Repository struct {
	User UserRepository
}

func Inject(db *gorm.DB) *Repository {
	return &Repository{User: NewUserRepository(db)}
}
