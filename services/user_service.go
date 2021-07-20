package services

import (
	"github.com/RamiAwar/go_demo_users_api/models/users"
	"github.com/RamiAwar/go_demo_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
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