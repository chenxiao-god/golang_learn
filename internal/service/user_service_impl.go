package service

import (
	"errors"
	"gomod/internal/module"
	"gomod/internal/repository"
)

type UserServiceImpl struct {
	repo *repository.UserRepository
}

func (u *UserServiceImpl) Create(user *module.User) error {
	if user.UserName == "" {
		// 返回错误，而不是仅仅创建错误
		return errors.New("username can not be empty")
	}
	// 调用仓储层的 Create 方法创建用户
	return (*u.repo).Create(user)
}

func (u *UserServiceImpl) Exit(username string) (*module.User, error) {
	return (*u.repo).FindByUsername(username)
}

func NewUserService(repo *repository.UserRepository) UserService {
	return &UserServiceImpl{repo: repo}
}
