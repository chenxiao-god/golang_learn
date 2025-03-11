package repository

import "gomod/internal/module"

type UserRepository interface {
	Create(user *module.User) error
	FindByUsername(username string) (*module.User, error)
}
