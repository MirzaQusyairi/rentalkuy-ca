package items

import (
	"rentalkuy-ca/business"
)

type ItemService struct {
	repository Repository
}

func NewItemService(repo Repository) Service {
	return &ItemService{
		repository: repo,
	}
}

func (serv *ItemService) Create(userID int, domain *Domain) (Domain, error) {

	result, err := serv.repository.Create(userID, domain)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	return result, nil
}
