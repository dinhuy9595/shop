package registry

import (
	ir "github.com/dinhuy9595/shop/infrastructure/repository"
	"github.com/dinhuy9595/shop/interface/api/controller"
	ip "github.com/dinhuy9595/shop/interface/presenter"
	"github.com/dinhuy9595/shop/usercases/interactor"
	up "github.com/dinhuy9595/shop/usercases/presenter"
	ur "github.com/dinhuy9595/shop/usercases/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
