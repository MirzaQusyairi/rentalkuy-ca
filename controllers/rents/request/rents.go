package request

import "rentalkuy-ca/business/rents"

type Rents struct {
	ItemID     int    `json:"item_id"`
	QTY        int    `json:"qty"`
	PacketID   int    `json:"packet_id"`
	Date       string `json:"date"`
	Duration   int    `json:"duration"`
	TotalQTY   int    `json:"total_qty"`
	TotalPrice int    `json:"total_price"`
	//Status     string `json:"status"`
}

func (req *Rents) ToDomain() *rents.Domain {
	return &rents.Domain{
		ItemID:     req.ItemID,
		QTY:        req.QTY,
		PacketID:   req.PacketID,
		Date:       req.Date,
		Duration:   req.Duration,
		TotalQTY:   req.TotalQTY,
		TotalPrice: req.TotalPrice,
		//Status:     req.Status,
	}
}
