package store

import "github.com/ankogit/http-rest-api/internal/app/model"

//UserRepository
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(string) (*model.User, error)
}