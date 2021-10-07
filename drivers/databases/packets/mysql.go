package packets

import (
	"rentalkuy-ca/business"
	"rentalkuy-ca/business/packets"

	"gorm.io/gorm"
)

type MysqlPacketRepository struct {
	Conn *gorm.DB
}

func NewMysqlPacketRepository(conn *gorm.DB) packets.Repository {
	return &MysqlPacketRepository{
		Conn: conn,
	}
}

func (rep *MysqlPacketRepository) Create(domain *packets.Domain) (packets.Domain, error) {
	photo := fromDomain(*domain)

	result := rep.Conn.Create(&photo)

	if result.Error != nil {
		return packets.Domain{}, result.Error
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

func (rep *MysqlPacketRepository) Delete(ID int) (string, error) {
	rec := Packets{}

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

func (rep *MysqlPacketRepository) GetByID(ID int) (packets.Domain, error) {
	var packet Packets

	result := rep.Conn.Where("id = ?", ID).First(&packet)

	if result.Error != nil {
		return packets.Domain{}, result.Error
	}

	return toDomain(packet), nil
}

func (rep *MysqlPacketRepository) GetAllByID(ID int) ([]packets.Domain, error) {
	var packet []Packets

	if err := rep.Conn.Where("item_id = ?", ID).Find(&packet).Error; err != nil {
		return []packets.Domain{}, err
	}

	return toDomainList(packet), nil

}
