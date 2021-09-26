package controller

import (
	"net/http"

	"github.com/dinhuy9595/shop/domain/model"
	"github.com/dinhuy9595/shop/usercases/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

// UserController is
type UserController interface {
	GetUsers(c Context) error
	GetStatus(c Context) error
}

// NewUserController is
func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, u)
	return nil
}
func (uc *userController) GetStatus(c Context) error {

	c.JSON(http.StatusOK, "okie")
	return nil
}
