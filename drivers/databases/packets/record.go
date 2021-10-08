package packets

import (
	"rentalkuy-ca/business/packets"
	"time"

	"gorm.io/gorm"
)

type Packets struct {
	gorm.Model
	ID        int `gorm:"primary_key"`
	ItemID    int
	Name      string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(packet Packets) packets.Domain {
	return packets.Domain{
		ID:        packet.ID,
		ItemID:    packet.ItemID,
		Name:      packet.Name,
		Price:     packet.Price,
		CreatedAt: packet.CreatedAt,
		UpdatedAt: packet.UpdatedAt,
	}
}

func fromDomain(domain packets.Domain) Packets {
	return Packets{
		ID:        domain.ID,
		ItemID:    domain.ItemID,
		Name:      domain.Name,
		Price:     domain.Price,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toDomainUpdate(packet Packets) packets.Domain {
	return packets.Domain{
		ID:        packet.ID,
		ItemID:    packet.ItemID,
		Name:      packet.Name,
		Price:     packet.Price,
		CreatedAt: packet.CreatedAt,
		UpdatedAt: packet.UpdatedAt,
	}
}

func toDomainList(data []Packets) []packets.Domain {
	result := []packets.Domain{}

	for _, val := range data {
		result = append(result, toDomain(val))
	}
	return result
}
