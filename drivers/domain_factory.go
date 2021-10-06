package drivers

import (
	userDomain "rentalkuy-ca/business/users"
	userDB "rentalkuy-ca/drivers/databases/users"

	itemDomain "rentalkuy-ca/business/items"
	itemDB "rentalkuy-ca/drivers/databases/items"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}

func NewItemRepository(conn *gorm.DB) itemDomain.Repository {
	return itemDB.NewMysqlItemRepository(conn)
}
