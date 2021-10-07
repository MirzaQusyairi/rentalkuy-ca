package rents

import (
	"rentalkuy-ca/business/rents"
	"time"

	"gorm.io/gorm"
)

type Rents struct {
	gorm.Model
	ID         int `gorm:"primary_key"`
	UserID     int
	ItemID     int
	QTY        int
	PacketID   int
	Date       time.Time
	Duration   int
	TotalQTY   int
	TotalPrice int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func toDomain(rent Rents) rents.Domain {
	return rents.Domain{
		ID:         rent.ID,
		UserID:     rent.UserID,
		ItemID:     rent.ItemID,
		QTY:        rent.QTY,
		PacketID:   rent.PacketID,
		Date:       rent.Date,
		Duration:   rent.Duration,
		TotalQTY:   rent.TotalQTY,
		TotalPrice: rent.TotalPrice,
		Status:     rent.Status,
		CreatedAt:  rent.CreatedAt,
		UpdatedAt:  rent.UpdatedAt,
	}
}

func fromDomain(domain rents.Domain) Rents {
	return Rents{
		ID:         domain.ID,
		UserID:     domain.UserID,
		ItemID:     domain.ItemID,
		QTY:        domain.QTY,
		PacketID:   domain.PacketID,
		Date:       domain.Date,
		Duration:   domain.Duration,
		TotalQTY:   domain.TotalQTY,
		TotalPrice: domain.TotalPrice,
		Status:     domain.Status,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func toDomainUpdate(rent Rents) rents.Domain {
	return rents.Domain{
		ID:         rent.ID,
		UserID:     rent.UserID,
		ItemID:     rent.ItemID,
		QTY:        rent.QTY,
		PacketID:   rent.PacketID,
		Date:       rent.Date,
		Duration:   rent.Duration,
		TotalQTY:   rent.TotalQTY,
		TotalPrice: rent.TotalPrice,
		Status:     rent.Status,
		CreatedAt:  rent.CreatedAt,
		UpdatedAt:  rent.UpdatedAt,
	}
}

func toDomainList(data []Rents) []rents.Domain {
	result := []rents.Domain{}

	for _, val := range data {
		result = append(result, toDomain(val))
	}
	return result
}
