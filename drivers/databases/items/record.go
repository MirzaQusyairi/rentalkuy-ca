package items

import (
	"rentalkuy-ca/business/items"
	"time"

	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	ID          int `gorm:"primary_key"`
	UserID      int
	Name        string
	CategoryID  int
	Description string
	QTY         int
	City        string
	Photo       string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func toDomain(item Items) items.Domain {
	return items.Domain{
		ID:          item.ID,
		UserID:      item.UserID,
		Name:        item.Name,
		CategoryID:  item.CategoryID,
		Description: item.Description,
		QTY:         item.QTY,
		City:        item.City,
		Photo:       item.Photo,
		Status:      item.Status,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func fromDomain(domain items.Domain) Items {
	return Items{
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

func toDomainUpdate(item Items) items.Domain {
	return items.Domain{
		ID:          item.ID,
		UserID:      item.UserID,
		Name:        item.Name,
		CategoryID:  item.CategoryID,
		Description: item.Description,
		QTY:         item.QTY,
		City:        item.City,
		Photo:       item.Photo,
		Status:      item.Status,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func toDomainList(data []Items) []items.Domain {
	result := []items.Domain{}

	for _, val := range data {
		result = append(result, toDomain(val))
	}
	return result
}
