package photos

import (
	"rentalkuy-ca/business"
	"rentalkuy-ca/business/photos"

	"gorm.io/gorm"
)

type MysqlPhotoRepository struct {
	Conn *gorm.DB
}

func NewMysqlPhotoRepository(conn *gorm.DB) photos.Repository {
	return &MysqlPhotoRepository{
		Conn: conn,
	}
}

func (rep *MysqlPhotoRepository) Create(domain *photos.Domain) (photos.Domain, error) {
	photo := fromDomain(*domain)

	result := rep.Conn.Create(&photo)

	if result.Error != nil {
		return photos.Domain{}, result.Error
	}

	return toDomain(photo), nil
}

// func (rep *MysqlPhotoRepository) Update(userID int, ID int, domain *items.Domain) (photos.Domain, error) {
// 	photoUpdate := fromDomain(*domain)

// 	find := rep.Conn.Where("id = ?", itemID).First(&itemUpdate).Error
// 	if find != nil {
// 		return items.Domain{}, business.ErrNotFound
// 	}

// 	itemUpdate.ID = itemID
// 	result := rep.Conn.Where("user_id = ?", userID).Where("id = ?", itemID).Updates(&itemUpdate)

// 	if result.Error != nil {
// 		return items.Domain{}, business.ErrNotFound
// 	}

// 	return toDomainUpdate(itemUpdate), nil
// }

func (rep *MysqlPhotoRepository) Delete(ID int) (string, error) {
	rec := Photos{}

	find := rep.Conn.Where("id = ?", ID).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", ID).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Photo has been delete", nil
}

func (rep *MysqlPhotoRepository) GetByID(ID int) (photos.Domain, error) {
	var photo Photos

	result := rep.Conn.Where("id = ?", ID).First(&photo)

	if result.Error != nil {
		return photos.Domain{}, result.Error
	}

	return toDomain(photo), nil
}

func (rep *MysqlPhotoRepository) GetAllByID(ID int) ([]photos.Domain, error) {
	var photo []Photos

	if err := rep.Conn.Where("item_id = ?", ID).Find(&photo).Error; err != nil {
		return []photos.Domain{}, err
	}

	return toDomainList(photo), nil

}
