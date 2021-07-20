package users

import (
	"fmt"

	"github.com/RamiAwar/go_demo_users_api/utils/errors"
)

// Fake database
var (
	usersDB = make(map[int64] *User)
)


func (user *User) Get() *errors.RestError {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}


func (user *User) Save() *errors.RestError {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("User with email %s is already registered", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}

	usersDB[user.Id] = user
	return nil
}