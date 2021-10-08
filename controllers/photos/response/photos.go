package response

import (
	"rentalkuy-ca/business/photos"
	"time"
)

type CreatePhotoResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id:"`
	ItemID    int       `json:"item_id:"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainCreatePhoto(domain photos.Domain) CreatePhotoResponse {
	return CreatePhotoResponse{
		Message:   "Create Photo Success",
		ID:        domain.ID,
		ItemID:    domain.ItemID,
		Path:      domain.Path,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

type PhotoResponse struct {
	ID        int       `json:"id:"`
	ItemID    int       `json:"item_id:"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomainPhoto(domain photos.Domain) PhotoResponse {
	return PhotoResponse{
		ID:        domain.ID,
		ItemID:    domain.ItemID,
		Path:      domain.Path,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainListItem(domain []photos.Domain) []PhotoResponse {
	var response []PhotoResponse
	for _, value := range domain {
		response = append(response, FromDomainPhoto(value))
	}
	return response
}
