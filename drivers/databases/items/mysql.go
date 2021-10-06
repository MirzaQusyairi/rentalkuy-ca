package items

import (
	"rentalkuy-ca/business/items"

	"gorm.io/gorm"
)

type MysqlItemRepository struct {
	Conn *gorm.DB
}

func NewMysqlItemRepository(conn *gorm.DB) items.Repository {
	return &MysqlItemRepository{
		Conn: conn,
	}
}

func (rep *MysqlItemRepository) Create(userID int, domain *items.Domain) (items.Domain, error) {
	item := fromDomain(*domain)
	item.UserID = userID
	result := rep.Conn.Create(&item)

	if result.Error != nil {
		return items.Domain{}, result.Error
	}

	return toDomain(item), nil

}
