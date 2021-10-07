package drivers

import (
	userDomain "rentalkuy-ca/business/users"
	userDB "rentalkuy-ca/drivers/databases/users"

	itemDomain "rentalkuy-ca/business/items"
	itemDB "rentalkuy-ca/drivers/databases/items"

	photoDomain "rentalkuy-ca/business/photos"
	photoDB "rentalkuy-ca/drivers/databases/photos"

	packetDomain "rentalkuy-ca/business/packets"
	packetDB "rentalkuy-ca/drivers/databases/packets"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}

func NewItemRepository(conn *gorm.DB) itemDomain.Repository {
	return itemDB.NewMysqlItemRepository(conn)
}

func NewPhotoRepository(conn *gorm.DB) photoDomain.Repository {
	return photoDB.NewMysqlPhotoRepository(conn)
}

func NewPacketRepository(conn *gorm.DB) packetDomain.Repository {
	return packetDB.NewMysqlPacketRepository(conn)
}
