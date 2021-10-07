package request

import "rentalkuy-ca/business/items"

type Items struct {
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"desc"`
	QTY         int    `json:"qty"`
	City        string `json:"city"`
	Photo       string `json:"photo"`
	Status      string `json:"status"`
}

func (req *Items) ToDomain() *items.Domain {
	return &items.Domain{
		Name:        req.Name,
		CategoryID:  req.CategoryID,
		Description: req.Description,
		QTY:         req.QTY,
		City:        req.City,
		Photo:       req.Photo,
		Status:      req.Status,
	}
}
