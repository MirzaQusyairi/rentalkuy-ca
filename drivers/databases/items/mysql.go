package items

import (
	"rentalkuy-ca/business"
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

func (rep *MysqlItemRepository) Update(userID int, itemID int, domain *items.Domain) (items.Domain, error) {
	itemUpdate := fromDomain(*domain)

	find := rep.Conn.Where("id = ?", itemID).First(&itemUpdate).Error
	if find != nil {
		return items.Domain{}, business.ErrNotFound
	}

	itemUpdate.ID = itemID
	result := rep.Conn.Where("user_id = ?", userID).Where("id = ?", itemID).Updates(&itemUpdate)

	if result.Error != nil {
		return items.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(itemUpdate), nil
}

func (rep *MysqlItemRepository) Delete(userID int, itemID int) (string, error) {
	rec := Items{}

	find := rep.Conn.Where("user_id = ?", userID).Where("id = ?", itemID).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", itemID).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Item has been delete", nil
}

func (rep *MysqlItemRepository) GetByID(itemID int) (items.Domain, error) {
	var item Items

	result := rep.Conn.Where("id = ?", itemID).First(&item)

	if result.Error != nil {
		return items.Domain{}, result.Error
	}

	return toDomain(item), nil
}

func (rep *MysqlItemRepository) GetAllByUserID(userID int) ([]items.Domain, error) {
	var item []Items

	if err := rep.Conn.Table("items").Joins("left join users on users.id = items.user_id").Where("user_id = ?", userID).Find(&item).Error; err != nil {
		return []items.Domain{}, err
	}

	// result := rep.Conn.Find(&item)

	// if result.Error != nil {
	// 	return []items.Domain{}, result.Error
	// }

	return toDomainList(item), nil

}
