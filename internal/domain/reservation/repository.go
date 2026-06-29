package reservation

import (
	"errors"

	"goproject/internal/domain/parkingzone"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	// CreateReservationTransaction(userID uint, zoneID uint, licensePlate string) (*Reservation, error)
CreateReservationTransaction(
		userID uint,
		zoneID uint,
		licensePlate string,
	) (*Reservation, error)
	GetByID(id uint) (*Reservation, error)

	Update(reservation *Reservation) error

	GetMyReservations(userID uint) ([]Reservation, error)

	GetAllReservations() ([]Reservation, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

// ========================================
// Create Reservation with Transaction
// FOR UPDATE Row Lock
// ========================================
// func (r *repository) CreateReservationTransaction(
// 	userID uint,
// 	zoneID uint,
// 	licensePlate string,
// ) (*Reservation, error) {

// 	var reservation Reservation

// 	err := r.db.Transaction(func(tx *gorm.DB) error {

// 		var zone parkingzone.ParkingZone

// 		// Lock parking zone row
// 		if err := tx.
// 			Clauses(clause.Locking{Strength: "UPDATE"}).
// 			First(&zone, zoneID).Error; err != nil {

// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				return errors.New("parking zone not found")
// 			}

// 			return err
// 		}

// 		// Count active reservations
// 		var activeCount int64

// 		if err := tx.
// 			Model(&Reservation{}).
// 			Where("zone_id = ? AND status = ?", zoneID, StatusActive).
// 			Count(&activeCount).Error; err != nil {
// 			return err
// 		}

// 		// Capacity check
// 		if activeCount >= int64(zone.TotalCapacity) {
// 			return errors.New("parking zone is full")
// 		}

// 		reservation = Reservation{
// 			UserID:       userID,
// 			ZoneID:       zoneID,
// 			LicensePlate: licensePlate,
// 			Status:       StatusActive,
// 		}

// 		if err := tx.Create(&reservation).Error; err != nil {
// 			return err
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &reservation, nil
// }


func (r *repository) CreateReservationTransaction(
	userID uint,
	zoneID uint,
	licensePlate string,
) (*Reservation, error) {

	var reservation Reservation

	err := r.db.Transaction(func(tx *gorm.DB) error {

		// 1. Lock parking zone row (FOR UPDATE)
		var zone parkingzone.ParkingZone

		if err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&zone, zoneID).Error; err != nil {

			return err
		}

		// 2. Count active reservations
		var activeCount int64

		if err := tx.
			Model(&Reservation{}).
			Where("zone_id = ? AND status = ?", zoneID, StatusActive).
			Count(&activeCount).Error; err != nil {

			return err
		}

		// 3. Capacity check
		if activeCount >= int64(zone.TotalCapacity) {
			return errors.New("parking zone is full")
		}

		// 4. Create reservation
		reservation = Reservation{
			UserID: userID,
			ZoneID: zoneID,

			LicensePlate: licensePlate,

			Status: StatusActive,
		}

		if err := tx.Create(&reservation).Error; err != nil {
			return err
		}

		// Load Zone and User (optional)
		if err := tx.
			Preload("User").
			Preload("Zone").
			First(&reservation, reservation.ID).Error; err != nil {

			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &reservation, nil
}


// ========================================
// Get Reservation By ID
// ========================================
func (r *repository) GetByID(id uint) (*Reservation, error) {

	var reservation Reservation

	err := r.db.
		Preload("User").
		Preload("Zone").
		First(&reservation, id).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &reservation, nil
}

// ========================================
// Update Reservation
// ========================================
func (r *repository) Update(reservation *Reservation) error {
	return r.db.Save(reservation).Error
}

// ========================================
// Get My Reservations
// ========================================
func (r *repository) GetMyReservations(userID uint) ([]Reservation, error) {

	var reservations []Reservation

	err := r.db.
		Preload("Zone").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&reservations).Error

	return reservations, err
}

// ========================================
// Get All Reservations (Admin)
// ========================================
func (r *repository) GetAllReservations() ([]Reservation, error) {

	var reservations []Reservation

	err := r.db.
		Preload("User").
		Preload("Zone").
		Order("created_at DESC").
		Find(&reservations).Error

	return reservations, err
}