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

func (rep *MysqlPacketRepository) Update(userID int, ID int, domain *packets.Domain) (packets.Domain, error) {
	packetUpdate := fromDomain(*domain)

	packetUpdate.ID = ID
	result := rep.Conn.Where("id = ?", ID).Updates(&packetUpdate)

	if result.Error != nil {
		return packets.Domain{}, business.ErrNotFound
	}

	return toDomainUpdate(packetUpdate), nil
}

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

	return "Packet has been delete", nil
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
