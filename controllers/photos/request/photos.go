package request

import (
	"rentalkuy-ca/business/photos"
)

type Photos struct {
	ItemID int    `json:"item_id"`
	Path   string `json:"path"`
}

func (req *Photos) ToDomain() *photos.Domain {
	return &photos.Domain{
		ItemID: req.ItemID,
		Path:   req.Path,
	}
}
