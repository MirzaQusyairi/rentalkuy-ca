package items

import (
	"time"
)

type Domain struct {
	ID          int
	UserID      int
	Name        string
	CategoryID  int
	Description string
	QTY         int
	City        string
	Photo       string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	Create(userID int, domain *Domain) (Domain, error)
	Update(userID int, itemID int, domain *Domain) (Domain, error)
	Delete(userID, itemID int) (string, error)
	GetByID(itemID int) (Domain, error)
	GetAllByUserID(userID int) ([]Domain, error)
}

type Repository interface {
	Create(userID int, domain *Domain) (Domain, error)
	Update(userID int, itemID int, domain *Domain) (Domain, error)
	Delete(userID, itemID int) (string, error)
	GetByID(itemID int) (Domain, error)
	GetAllByUserID(userID int) ([]Domain, error)
}
