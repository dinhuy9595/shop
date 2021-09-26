package interactor

import (
	"github.com/dinhuy9595/shop/domain/model"
	"github.com/dinhuy9595/shop/usercases/presenter"
	"github.com/dinhuy9595/shop/usercases/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

//UserInteractor is
type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
}

// NewUserInteractor is
func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}
