package response

import (
	"rentalkuy-ca/business/items"
	"time"
)

type CreateItemResponse struct {
	Message     string    `json:"message"`
	ID          int       `json:"id:"`
	Name        string    `json:"name"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"desc"`
	QTY         int       `json:"qty"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomainCreateItem(domain items.Domain) CreateItemResponse {
	return CreateItemResponse{
		Message:     "Create Item Success",
		ID:          domain.ID,
		Name:        domain.Name,
		CategoryID:  domain.CategoryID,
		Description: domain.Description,
		QTY:         domain.QTY,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
