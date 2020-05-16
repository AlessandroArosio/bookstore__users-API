package services

import (
	"github.com/alessandroarosio/bookstore_users-API/domain/users"
	"github.com/alessandroarosio/bookstore_users-API/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
