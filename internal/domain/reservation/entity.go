package reservation

import (
	"time"

	"goproject/internal/domain/parkingzone"
	"goproject/internal/domain/user"

	"gorm.io/gorm"
)

const (
	StatusActive    = "active"
	StatusCompleted = "completed"
	StatusCancelled = "cancelled"
)

type Reservation struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint      `gorm:"not null" json:"user_id"`
	User   user.User `gorm:"foreignKey:UserID" json:"user"`

	ZoneID uint                    `gorm:"not null" json:"zone_id"`
	Zone   parkingzone.ParkingZone `gorm:"foreignKey:ZoneID" json:"zone"`

	LicensePlate string `gorm:"size:15;not null" json:"license_plate"`

	Status string `gorm:"type:varchar(20);default:'active';not null" json:"status"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}