package response

import (
	"rentalkuy-ca/business/items"
	"time"
)

type CreateItemResponse struct {
	Message     string    `json:"message"`
	ID          int       `json:"id:"`
	UserID      int       `json:"user_id:"`
	Name        string    `json:"name"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"desc"`
	QTY         int       `json:"qty"`
	City        string    `json:"city"`
	Photo       string    `json:"photo"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomainCreateItem(domain items.Domain) CreateItemResponse {
	return CreateItemResponse{
		Message:     "Create Item Success",
		ID:          domain.ID,
		UserID:      domain.UserID,
		Name:        domain.Name,
		CategoryID:  domain.CategoryID,
		Description: domain.Description,
		QTY:         domain.QTY,
		City:        domain.City,
		Photo:       domain.Photo,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

type ItemResponse struct {
	ID          int       `json:"id:"`
	UserID      int       `json:"user_id:"`
	Name        string    `json:"name"`
	CategoryID  int       `json:"category_id"`
	Description string    `json:"desc"`
	QTY         int       `json:"qty"`
	City        string    `json:"city"`
	Photo       string    `json:"photo"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomainItem(domain items.Domain) ItemResponse {
	return ItemResponse{
		ID:          domain.ID,
		Name:        domain.Name,
		UserID:      domain.UserID,
		CategoryID:  domain.CategoryID,
		Description: domain.Description,
		QTY:         domain.QTY,
		City:        domain.City,
		Photo:       domain.Photo,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromDomainUpdateItem(domain items.Domain) CreateItemResponse {
	return CreateItemResponse{
		Message:     "Update Event Success",
		ID:          domain.ID,
		Name:        domain.Name,
		UserID:      domain.UserID,
		CategoryID:  domain.CategoryID,
		Description: domain.Description,
		QTY:         domain.QTY,
		City:        domain.City,
		Photo:       domain.Photo,
		Status:      domain.Status,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func FromDomainListItem(domain []items.Domain) []ItemResponse {
	var response []ItemResponse
	for _, value := range domain {
		response = append(response, FromDomainItem(value))
	}
	return response
}
