package services

import (
	"github.com/RamiAwar/go_demo_users_api/models/user"
	"github.com/RamiAwar/go_demo_users_api/utils/errors"
)

func CreateUser(user user.User) (*user.User, *errors.RestError) {
	return &user, nil
}