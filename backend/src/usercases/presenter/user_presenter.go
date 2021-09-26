package presenter

import "github.com/dinhuy9595/shop/domain/model"

//UserPresenter is
type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
