package request

import (
	"rentalkuy-ca/business/users"
)

// request body for login
type UsersLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// request body for register
type Users struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Location string `json:"location"`
}

// turn register body to domain
func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Username: req.Username,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		Address:  req.Address,
		Location: req.Location,
	}
}
