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
	GetAllByUserID(userID int) ([]Domain, error)
	// GetByID(userID, id int) (Domain, error)
	// DeleteItems(userID int, itemsID int)
}

type Repository interface {
	Create(userID int, domain *Domain) (Domain, error)
	Update(userID int, itemID int, domain *Domain) (Domain, error)
	Delete(userID, itemID int) (string, error)
	GetAllByUserID(userID int) ([]Domain, error)
	// GetByID(userID, watchlistId int) (Domain, error)
	// DeleteItems(userID int, itemsID int)
}
