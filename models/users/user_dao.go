package users

import (
	"fmt"
	"strings"

	"github.com/RamiAwar/go_demo_users_api/utils/date"
	"github.com/RamiAwar/go_demo_users_api/utils/errors"
	"github.com/RamiAwar/go_demo_users_api/utils/security"
)

// Fake database
var (
	usersDB = make(map[int64]*User)
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

	user.DateCreated = date.Now()
	hash, err := security.HashPassword(user.Password)
	if err != nil {
		return errors.InternalServerError("Could not save user, contact support.")
	}

	user.Password = hash
	usersDB[user.Id] = user
	return nil
}

func Find(query string) (Users, *errors.RestError) {
	var result Users

	for _, v := range usersDB {
		user := User{
			FirstName: strings.ToLower(v.FirstName),
			LastName:  strings.ToLower(v.LastName),
			Email:     strings.ToLower(v.Email),
		}

		query = strings.ToLower(query)

		if strings.Contains(user.FirstName, query) ||
			strings.Contains(user.LastName, query) ||
			strings.Contains(user.Email, query) ||
			strings.Contains(strings.Join([]string{user.FirstName, user.LastName}, " "), query) {
			result = append(result, *v)
		}
	}

	return result, nil
}
