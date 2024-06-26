package store

import "github.com/drummdaddy/http-rest-api/cmd/apiserver/internal/app/model"

type UserRepository interface {
	Create(*model.User) error

	FindByEmail(string) (*model.User, error)
}
