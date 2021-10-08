package photos

import (
	"rentalkuy-ca/business/photos"
	"time"

	"gorm.io/gorm"
)

type Photos struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	ItemID    int
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(photo Photos) photos.Domain {
	return photos.Domain{
		ID:        photo.ID,
		ItemID:    photo.ItemID,
		Path:      photo.Path,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	}
}

func fromDomain(domain photos.Domain) Photos {
	return Photos{
		ID:        domain.ID,
		ItemID:    domain.ItemID,
		Path:      domain.Path,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toDomainList(data []Photos) []photos.Domain {
	result := []photos.Domain{}

	for _, val := range data {
		result = append(result, toDomain(val))
	}
	return result
}
