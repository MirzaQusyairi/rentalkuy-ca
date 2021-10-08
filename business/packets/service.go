package packets

import (
	"rentalkuy-ca/business"
	"rentalkuy-ca/business/users"
)

type PacketService struct {
	repository     Repository
	userRepository users.Repository
}

func NewPacketService(repo Repository, userRepo users.Repository) Service {
	return &PacketService{
		repository:     repo,
		userRepository: userRepo,
	}
}

func (serv *PacketService) Create(domain *Domain) (Domain, error) {
	result, err := serv.repository.Create(domain)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

// func (serv *PhotoService) Update(userID int, ID int, domain *Domain) (Domain, error) {
// 	result, err := serv.repository.Update(userID, ID, domain)

// 	if err != nil {
// 		return Domain{}, err
// 	}

// 	return result, nil
// }

func (serv *PacketService) Delete(ID int) (string, error) {

	result, err := serv.repository.Delete(ID)

	if err != nil {
		return "", business.ErrNotFound
	}

	return result, nil
}

func (serv *PacketService) GetByID(ID int) (Domain, error) {
	result, err := serv.repository.GetByID(ID)

	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (serv *PacketService) GetAllByID(ID int) ([]Domain, error) {
	result, _ := serv.repository.GetAllByID(ID)

	if result == nil {
		return nil, business.ErrNotFound
	}

	return result, nil
}
