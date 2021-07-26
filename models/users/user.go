package users

import (
	"encoding/json"
	"strings"

	"github.com/RamiAwar/go_demo_users_api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateCreated string `json:"date_created"`
}

type Users []User

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestError("Email is required")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.BadRequestError("Password should not be empty")
	}
	return nil
}

type UserOut struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Marshal() interface{} {
	userJson, _ := json.Marshal(user)
	var out UserOut
	if err := json.Unmarshal(userJson, &out); err != nil {
		return nil
	}
	return out
}

func (users Users) Marshal() []interface{} {
	result := make([]interface{}, len(users))
	for i, user := range users {
		result[i] = user.Marshal()
	}
	return result
}
