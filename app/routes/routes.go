package routes

import (
	//middlewareApp "rentalkuy-ca/app/middlewares"
	"rentalkuy-ca/controllers/items"
	"rentalkuy-ca/controllers/packets"
	"rentalkuy-ca/controllers/photos"
	"rentalkuy-ca/controllers/rents"
	"rentalkuy-ca/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware    middleware.JWTConfig
	UserController   users.UserController
	ItemController   items.ItemController
	PhotoController  photos.PhotoController
	PacketController packets.PacketController
	RentController   rents.RentController
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
	items.GET("/available", cl.ItemController.GetAll) //public

	//photo
	items.POST("/photo/insert", cl.PhotoController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.DELETE("/photo/delete/:id", cl.PhotoController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.GET("/photo/:id", cl.PhotoController.GetByID)        //public
	items.GET("/photo/all/:id", cl.PhotoController.GetAllByID) //public

	//packet
	items.POST("/packet/insert", cl.PacketController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.PUT("/packet/update/:id", cl.PacketController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.DELETE("/packet/delete/:id", cl.PacketController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))
	items.GET("/packet/:id", cl.PacketController.GetByID)        //public
	items.GET("/packet/all/:id", cl.PacketController.GetAllByID) //public

	//rent
	rents := e.Group("rents")
	rents.POST("/insert", cl.RentController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
	rents.DELETE("/delete/:id", cl.RentController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware))
	rents.GET("/:id", cl.RentController.GetByID, middleware.JWTWithConfig(cl.JWTMiddleware))
	rents.GET("/all", cl.RentController.GetAllByUserID, middleware.JWTWithConfig(cl.JWTMiddleware))
}
