package routes

import (
	//middlewareApp "go-watchlist/app/middlewares"
	"rentalkuy-ca/controllers/items"
	"rentalkuy-ca/controllers/photos"
	"rentalkuy-ca/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware   middleware.JWTConfig
	UserController  users.UserController
	ItemController  items.ItemController
	PhotoController photos.PhotoController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	//users
	users := e.Group("users")
	users.POST("/register", cl.UserController.Register)
	users.POST("/login", cl.UserController.Login)

	//items
	items := e.Group("items")
	items.POST("/insert", cl.ItemController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.PUT("/update/:id", cl.ItemController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.DELETE("/delete/:id", cl.ItemController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.GET("/:id", cl.ItemController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.GET("/all", cl.ItemController.GetAllByUserID, middleware.JWTWithConfig(cl.JWTMiddleware))

	items.POST("/photo/insert", cl.PhotoController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.DELETE("/photo/delete/:id", cl.PhotoController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.GET("/photo/:id", cl.PhotoController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.GET("/photo/all/:id", cl.PhotoController.GetAllByID, middleware.JWTWithConfig(cl.JWTMiddleware))
}
