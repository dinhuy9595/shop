package router

import (
	"github.com/dinhuy9595/shop/interface/api/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//NewRouter is
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error {
		return c.User.GetUsers(context)

	})
	e.GET("/sttus", func(context echo.Context) error {
		return c.User.GetStatus(context)

	})

	return e
}
