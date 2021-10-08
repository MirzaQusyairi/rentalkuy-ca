package request

import "rentalkuy-ca/business/packets"

type Packets struct {
	ItemID int    `json:"item_id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}

func (req *Packets) ToDomain() *packets.Domain {
	return &packets.Domain{
		ItemID: req.ItemID,
		Name:   req.Name,
		Price:  req.Price,
	}
}
