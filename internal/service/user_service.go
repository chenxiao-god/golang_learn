package service

import "gomod/internal/module"

type UserService interface {
	Create(user *module.User) error
	Exit(username string) (*module.User, error)
}
