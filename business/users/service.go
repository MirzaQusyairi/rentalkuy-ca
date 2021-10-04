package users

import (
	"time"

	"rentalkuy-ca/app/middlewares"
	"rentalkuy-ca/business"
	"rentalkuy-ca/helper/encrypt"
)

type UserService struct {
	repository     Repository
	contextTimeout time.Duration
	jwtAuth        *middlewares.ConfigJWT
}

func NewUserService(repo Repository, timeout time.Duration, jwtauth *middlewares.ConfigJWT) Service {
	return &UserService{
		repository:     repo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

func (servUser *UserService) Login(username, password string) (Domain, error) {
	if username == "" {
		return Domain{}, business.ErrUsernameEmpty
	}
	if password == "" {
		return Domain{}, business.ErrPasswordEmpty
	}

	user, err := servUser.repository.Login(username, password)
	if err != nil {
		return Domain{}, err
	}
	validPass := encrypt.ValidateHash(password, user.Password)
	if !validPass {
		return Domain{}, business.ErrWrongPassword
	}
	if err != nil {
		return Domain{}, err
	}
	user.Token = servUser.jwtAuth.GenerateToken(user.ID, "user")
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (servUser *UserService) Register(domain *Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, business.ErrUsernameEmpty
	}
	if domain.Email == "" {
		return Domain{}, business.ErrEmailEmpty
	}
	if domain.Password == "" {
		return Domain{}, business.ErrPasswordEmpty
	}
	encryptedPass, _ := encrypt.HashPassword(domain.Password)
	domain.Password = encryptedPass
	user, err := servUser.repository.Register(domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
