package presenter

import "github.com/dinhuy9595/shop/domain/model"

type userPresenter struct {
}

// UserPresenter is
type UserPresenter interface {
	ResponseUsers(us []*model.User) []*model.User
}

// NewUserPresenter is
func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}
