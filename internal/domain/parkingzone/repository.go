package parkingzone

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(zone *ParkingZone) error
	GetAll() ([]ParkingZone, error)
	GetByID(id uint) (*ParkingZone, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// Create a parking zone
func (r *repository) Create(zone *ParkingZone) error {
	return r.db.Create(zone).Error
}

// Get all parking zones
func (r *repository) GetAll() ([]ParkingZone, error) {
	var zones []ParkingZone

	err := r.db.
		Order("id ASC").
		Find(&zones).Error

	if err != nil {
		return nil, err
	}

	return zones, nil
}

// Get a parking zone by ID
func (r *repository) GetByID(id uint) (*ParkingZone, error) {
	var zone ParkingZone

	err := r.db.First(&zone, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &zone, nil
}