package rents

import (
	"rentalkuy-ca/business"
	"rentalkuy-ca/business/rents"

	"gorm.io/gorm"
)

type MysqlRentRepository struct {
	Conn *gorm.DB
}

func NewMysqlRentRepository(conn *gorm.DB) rents.Repository {
	return &MysqlRentRepository{
		Conn: conn,
	}
}

func (rep *MysqlRentRepository) Create(userID int, domain *rents.Domain) (rents.Domain, error) {
	rent := fromDomain(*domain)
	rent.UserID = userID
	rent.Status = "waiting"
	result := rep.Conn.Create(&rent)

	if result.Error != nil {
		return rents.Domain{}, result.Error
	}

	return toDomain(rent), nil
}

func (rep *MysqlRentRepository) Delete(userID int, rentID int) (string, error) {
	rec := Rents{}

	find := rep.Conn.Where("user_id = ?", userID).Where("id = ?", rentID).First(&rec)

	if find.Error != nil {
		return "", business.ErrUnathorized
	}

	err := rep.Conn.Delete(&rec, "id = ?", rentID).Error

	if err != nil {
		return "", business.ErrNotFound
	}

	return "Rent has been delete", nil
}

func (rep *MysqlRentRepository) GetByID(rentID int) (rents.Domain, error) {
	var rent Rents

	result := rep.Conn.Where("id = ?", rentID).First(&rent)

	if result.Error != nil {
		return rents.Domain{}, result.Error
	}

	return toDomain(rent), nil
}

func (rep *MysqlRentRepository) GetAllByUserID(userID int) ([]rents.Domain, error) {
	var rent []Rents

	if err := rep.Conn.Table("rents").Joins("left join users on users.id = rents.user_id").Where("user_id = ?", userID).Find(&rent).Error; err != nil {
		return []rents.Domain{}, err
	}

	return toDomainList(rent), nil
}
