package items

import (
	"rentalkuy-ca/business"
	"rentalkuy-ca/business/geolocation"
	"rentalkuy-ca/business/users"
)

type ItemService struct {
	repository     Repository
	userRepository users.Repository
	geoRepository  geolocation.Repository
}

func NewItemService(repo Repository, userRepo users.Repository, geoRepo geolocation.Repository) Service {
	return &ItemService{
		repository:     repo,
		userRepository: userRepo,
		geoRepository:  geoRepo,
	}
}

func (serv *ItemService) Create(userID int, ip string, domain *Domain) (Domain, error) {
	location, err := serv.geoRepository.GetLocationByIP(ip)
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	domain.City = location.City

	result, err := serv.repository.Create(userID, ip, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *ItemService) Update(userID int, itemID int, domain *Domain) (Domain, error) {
	result, err := serv.repository.Update(userID, itemID, domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *ItemService) Delete(userID int, itemID int) (string, error) {
	result, err := serv.repository.Delete(userID, itemID)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}

func (serv *ItemService) GetByID(itemID int) (Domain, error) {
	result, err := serv.repository.GetByID(itemID)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *ItemService) GetAllByUserID(userID int) ([]Domain, error) {
	result, _ := serv.repository.GetAllByUserID(userID)

	if result == nil {
		return nil, business.ErrNotFound
	}

	return result, nil
}

func (serv *ItemService) GetAll() ([]Domain, error) {
	result, _ := serv.repository.GetAll()

	if result == nil {
		return nil, business.ErrNotFound
	}

	return result, nil
}
