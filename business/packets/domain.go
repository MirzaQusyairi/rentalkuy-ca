package packets

import (
	"time"
)

type Domain struct {
	ID        int
	ItemID    int
	Name      string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Create(domain *Domain) (Domain, error)
	//Update(userID int, ID int, domain *Domain) (Domain, error)
	Delete(ID int) (string, error)
	GetByID(ID int) (Domain, error)
	GetAllByID(ID int) ([]Domain, error)
}

type Repository interface {
	Create(domain *Domain) (Domain, error)
	//Update(userID int, ID int, domain *Domain) (Domain, error)
	Delete(ID int) (string, error)
	GetByID(ID int) (Domain, error)
	GetAllByID(ID int) ([]Domain, error)
}
