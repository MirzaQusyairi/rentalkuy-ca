package routes

import (
	// middlewareApp "go-watchlist/app/middlewares"
	"rentalkuy-ca/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Register)
	users.POST("/login", cl.UserController.Login)
}
