package router

import (
	"net/http"

	"github.com/dinhuy9595/shop/domain/common/errorHandle"
	"github.com/dinhuy9595/shop/interface/api/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//NewRouter is
func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		var (
			code = http.StatusInternalServerError
			key  = "ServerError"
			msg  string
		)

		if he, ok := err.(*errorHandle.HttpError); ok {
			code = he.Code
			key = he.Key
			msg = he.Message
		} else {
			msg = http.StatusText(code)
		}

		c.JSON(code, errorHandle.NewHTTPError(code, key, msg))

	}
	e.GET("/users", func(context echo.Context) error {
		return c.User.GetUsers(context)

	})
	e.GET("/sttus", func(context echo.Context) error {
		return c.User.GetStatus(context)

	})

	return e
}
