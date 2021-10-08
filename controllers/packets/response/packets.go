package response

import (
	"rentalkuy-ca/business/packets"
	"time"
)

type CreatePacketResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	ItemID    int       `json:"item_id:"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainCreatePacket(domain packets.Domain) CreatePacketResponse {
	return CreatePacketResponse{
		Message:   "Create Packet Success",
		ID:        domain.ID,
		ItemID:    domain.ItemID,
		Name:      domain.Name,
		Price:     domain.Price,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type PacketResponse struct {
	ID        int       `json:"id:"`
	ItemID    int       `json:"item_id:"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainPacket(domain packets.Domain) PacketResponse {
	return PacketResponse{
		ID:        domain.ID,
		ItemID:    domain.ItemID,
		Name:      domain.Name,
		Price:     domain.Price,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

// func FromDomainUpdatePhoto(domain items.Domain) CreateItemResponse {
// 	return CreateItemResponse{
// 		Message:     "Update Event Success",
// 		ID:          domain.ID,
// 		Name:        domain.Name,
// 		UserID:      domain.UserID,
// 		CategoryID:  domain.CategoryID,
// 		Description: domain.Description,
// 		QTY:         domain.QTY,
// 		City:        domain.City,
// 		Photo:       domain.Photo,
// 		Status:      domain.Status,
// 		CreatedAt:   domain.CreatedAt,
// 		UpdatedAt:   domain.UpdatedAt,
// 	}
// }

func FromDomainListPacket(domain []packets.Domain) []PacketResponse {
	var response []PacketResponse
	for _, value := range domain {
		response = append(response, FromDomainPacket(value))
	}
	return response
}
