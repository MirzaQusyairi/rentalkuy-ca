package users

import (
	"time"
)

type Domain struct {
	ID        int
	Username  string
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	Location  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Register(domain *Domain) (Domain, error)
	Login(username string, password string) (Domain, error)
}

type Repository interface {
	Register(domain *Domain) (Domain, error)
	Login(username string, password string) (Domain, error)
}
