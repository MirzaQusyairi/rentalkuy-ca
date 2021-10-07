package main

import (
	"log"

	_driverFactory "rentalkuy-ca/drivers"

	_userService "rentalkuy-ca/business/users"
	_userController "rentalkuy-ca/controllers/users"
	_userRepo "rentalkuy-ca/drivers/databases/users"

	_itemService "rentalkuy-ca/business/items"
	_itemController "rentalkuy-ca/controllers/items"
	_itemRepo "rentalkuy-ca/drivers/databases/items"

	_photoService "rentalkuy-ca/business/photos"
	_photoController "rentalkuy-ca/controllers/photos"
	_photoRepo "rentalkuy-ca/drivers/databases/photos"

	_packetService "rentalkuy-ca/business/packets"
	_packetController "rentalkuy-ca/controllers/packets"
	_packetRepo "rentalkuy-ca/drivers/databases/packets"

	_rentService "rentalkuy-ca/business/rents"
	_rentController "rentalkuy-ca/controllers/rents"
	_rentRepo "rentalkuy-ca/drivers/databases/rents"

	_dbDriver "rentalkuy-ca/drivers/mysql"

	_middleware "rentalkuy-ca/app/middlewares"
	_routes "rentalkuy-ca/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.Users{},
		&_itemRepo.Items{},
		&_photoRepo.Photos{},
		&_packetRepo.Packets{},
		&_rentRepo.Rents{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDB.InitDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: int64(viper.GetInt(`jwt.expired`)),
	}

	e := echo.New()
	e.IPExtractor = echo.ExtractIPDirect()
	geoRepo := _driverFactory.NewGeolocationRepository()

	userRepo := _driverFactory.NewUserRepository(db)
	userService := _userService.NewUserService(userRepo, 10, &configJWT)
	userCtrl := _userController.NewUserController(userService)

	itemRepo := _driverFactory.NewItemRepository(db)
	itemService := _itemService.NewItemService(itemRepo, userRepo, geoRepo)
	itemCtrl := _itemController.NewItemController(itemService)

	photoRepo := _driverFactory.NewPhotoRepository(db)
	photoService := _photoService.NewPhotoService(photoRepo, userRepo)
	photoCtrl := _photoController.NewPhotoController(photoService)

	packetRepo := _driverFactory.NewPacketRepository(db)
	packetService := _packetService.NewPacketService(packetRepo, userRepo)
	packetCtrl := _packetController.NewPacketController(packetService)

	rentRepo := _driverFactory.NewRentRepository(db)
	rentService := _rentService.NewRentService(rentRepo, userRepo)
	rentCtrl := _rentController.NewRentController(rentService)

	routesInit := _routes.ControllerList{
		JWTMiddleware:    configJWT.Init(),
		UserController:   *userCtrl,
		ItemController:   *itemCtrl,
		PhotoController:  *photoCtrl,
		PacketController: *packetCtrl,
		RentController:   *rentCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
