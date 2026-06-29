package parkingzone

import "gorm.io/gorm"

type ParkingZone struct {
	gorm.Model

	Name          string  `json:"name" gorm:"type:varchar(150);not null"`
	Type          string  `json:"type" gorm:"type:varchar(30);not null"`
	TotalCapacity int     `json:"total_capacity" gorm:"not null;check:total_capacity > 0"`
	PricePerHour  float64 `json:"price_per_hour" gorm:"type:decimal(10,2);not null;check:price_per_hour > 0"`
}