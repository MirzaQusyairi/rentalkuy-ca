package users

import (
	"rentalkuy-ca/business/users"
	"time"
)

type Users struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Phone     string
	Address   string
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(user Users) users.Domain {
	return users.Domain{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
		Address:   user.Address,
		Location:  user.Location,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func fromDomain(domain users.Domain) Users {
	return Users{
		ID:        domain.ID,
		Username:  domain.Username,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		Phone:     domain.Phone,
		Address:   domain.Address,
		Location:  domain.Location,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
