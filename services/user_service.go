package services

import (
	"github.com/JAlayon/bookstore_users-api/domain/users"
	"github.com/JAlayon/bookstore_users-api/errors"
)

func CreateUser(user users.User) (*users.User, *errors.ApiError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.ApiError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
