package rents

import (
	"time"
)

type Domain struct {
	ID         int
	UserID     int
	ItemID     int
	QTY        int
	PacketID   int
	Date       string
	Duration   int
	TotalQTY   int
	TotalPrice int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	Create(userID int, domain *Domain) (Domain, error)
	Delete(userID, itemID int) (string, error)
	GetByID(itemID int) (Domain, error)
	GetAllByUserID(userID int) ([]Domain, error)
}

type Repository interface {
	Create(userID int, domain *Domain) (Domain, error)
	Delete(userID, itemID int) (string, error)
	GetByID(itemID int) (Domain, error)
	GetAllByUserID(userID int) ([]Domain, error)
}
