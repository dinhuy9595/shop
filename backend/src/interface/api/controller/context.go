package controller

//Context is
type Context interface {
	JSON(code int, i interface{}) error
}
