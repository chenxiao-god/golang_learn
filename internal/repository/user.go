package repository

import (
	"gomod/internal/database"
	"gomod/internal/module"
)

type userRepository struct {
	db *database.DB
}

func (u *userRepository) Create(user *module.User) error {

	return u.db.Create(user).Error
}

func (u *userRepository) FindByUsername(username string) (*module.User, error) {
	var user module.User
	err := u.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
func NewUserRepository(db *database.DB) UserRepository {
	return &userRepository{db: db}
}
