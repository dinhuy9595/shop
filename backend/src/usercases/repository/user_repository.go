package repository

import "github.com/dinhuy9595/shop/domain/model"

//UserRepository is repo of user
type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
