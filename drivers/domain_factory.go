package drivers

import (
	userDomain "rentalkuy-ca/business/users"
	userDB "rentalkuy-ca/drivers/databases/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}
