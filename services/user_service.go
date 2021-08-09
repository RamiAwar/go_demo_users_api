package services

import (
	"github.com/RamiAwar/go_demo_users_api/models/users"
	"github.com/RamiAwar/go_demo_users_api/utils/errors"
)

var (
	UserService IUserService = &userService{}
)

type userService struct {
}

type IUserService interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(int64) (*users.User, *errors.RestError)
	Search(string) (users.Users, *errors.RestError)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestError) {
	if userId <= 0 {
		return nil, errors.BadRequestError("Invalid user ID")
	}

	result := &users.User{Id: userId}

	// Note automatic dereferencing below, cool feature
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *userService) Search(query string) (users.Users, *errors.RestError) {
	return users.Find(query)
}
