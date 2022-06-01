package users

import (
	"fmt"

	"github.com/JAlayon/bookstore_users-api/datasources"
	"github.com/JAlayon/bookstore_users-api/errors"
	"github.com/JAlayon/bookstore_users-api/utils"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Save() *errors.ApiError {
	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	user.DateCreated = utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}

func (user *User) Get() *errors.ApiError {
	err := datasources.Client.Ping()
	if err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
