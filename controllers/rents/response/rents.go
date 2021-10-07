package response

import (
	"rentalkuy-ca/business/rents"
	"time"
)

type CreateRentResponse struct {
	Message    string    `json:"message"`
	ID         int       `json:"id:"`
	UserID     int       `json:"user_id:"`
	ItemID     int       `json:"item_id"`
	QTY        int       `json:"qty"`
	PacketID   int       `json:"packet_id"`
	Duration   int       `json:"duration"`
	TotalQTY   int       `json:"total_qty"`
	TotalPrice int       `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomainCreateRent(domain rents.Domain) CreateRentResponse {
	return CreateRentResponse{
		Message:    "Create Rent Success",
		ID:         domain.ID,
		UserID:     domain.UserID,
		ItemID:     domain.ItemID,
		QTY:        domain.QTY,
		PacketID:   domain.PacketID,
		Duration:   domain.Duration,
		TotalQTY:   domain.TotalQTY,
		TotalPrice: domain.TotalPrice,
		Status:     domain.Status,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

type RentResponse struct {
	ID         int       `json:"id:"`
	UserID     int       `json:"user_id:"`
	ItemID     int       `json:"item_id"`
	QTY        int       `json:"qty"`
	PacketID   int       `json:"packet_id"`
	Duration   int       `json:"duration"`
	TotalQTY   int       `json:"total_qty"`
	TotalPrice int       `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomainItem(domain rents.Domain) RentResponse {
	return RentResponse{
		ID:         domain.ID,
		UserID:     domain.UserID,
		ItemID:     domain.ItemID,
		QTY:        domain.QTY,
		PacketID:   domain.PacketID,
		Duration:   domain.Duration,
		TotalQTY:   domain.TotalQTY,
		TotalPrice: domain.TotalPrice,
		Status:     domain.Status,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromDomainUpdateItem(domain rents.Domain) CreateRentResponse {
	return CreateRentResponse{
		Message:    "Update Rent Success",
		ID:         domain.ID,
		UserID:     domain.UserID,
		ItemID:     domain.ItemID,
		QTY:        domain.QTY,
		PacketID:   domain.PacketID,
		Duration:   domain.Duration,
		TotalQTY:   domain.TotalQTY,
		TotalPrice: domain.TotalPrice,
		Status:     domain.Status,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromDomainListItem(domain []rents.Domain) []RentResponse {
	var response []RentResponse
	for _, value := range domain {
		response = append(response, FromDomainItem(value))
	}
	return response
}
