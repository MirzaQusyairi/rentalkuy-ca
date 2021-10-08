package rents

import (
	"rentalkuy-ca/business"
	"rentalkuy-ca/business/users"
)

type RentService struct {
	repository     Repository
	userRepository users.Repository
}

func NewRentService(repo Repository, userRepo users.Repository) Service {
	return &RentService{
		repository:     repo,
		userRepository: userRepo,
	}
}

func (serv *RentService) Create(userID int, domain *Domain) (Domain, error) {
	result, err := serv.repository.Create(userID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *RentService) Delete(userID int, rentID int) (string, error) {
	result, err := serv.repository.Delete(userID, rentID)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}

func (serv *RentService) GetByID(rentID int) (Domain, error) {
	result, err := serv.repository.GetByID(rentID)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *RentService) GetAllByUserID(userID int) ([]Domain, error) {
	result, _ := serv.repository.GetAllByUserID(userID)

	if result == nil {
		return nil, business.ErrNotFound
	}

	return result, nil
}
