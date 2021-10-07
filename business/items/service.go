package items

import (
	"rentalkuy-ca/business"
	"rentalkuy-ca/business/users"
)

type ItemService struct {
	repository     Repository
	userRepository users.Repository
}

func NewItemService(repo Repository, userRepo users.Repository) Service {
	return &ItemService{
		repository:     repo,
		userRepository: userRepo,
	}
}

func (serv *ItemService) Create(userID int, domain *Domain) (Domain, error) {
	result, err := serv.repository.Create(userID, domain)

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

	// for i, val := range result {
	// 	result[i].UserName = serv.userRepository.GetNameByUserID(val.UserID)
	// }

	return result, nil
}
