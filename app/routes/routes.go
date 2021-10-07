package routes

import (
	//middlewareApp "go-watchlist/app/middlewares"
	"rentalkuy-ca/controllers/items"
	"rentalkuy-ca/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
	ItemController items.ItemController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Register)
	users.POST("/login", cl.UserController.Login)
	users.POST("/items/insert", cl.ItemController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	users.GET("/items/all", cl.ItemController.GetAllByUserID, middleware.JWTWithConfig(cl.JWTMiddleware))
	users.PUT("/items/update/:id", cl.ItemController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	users.DELETE("/items/delete/:id", cl.ItemController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))
}
