package reservation

import (
	"errors"
	"fmt"

	"goproject/internal/domain/reservation/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

// ==========================================
// Reserve Parking Spot
// ==========================================
func (s *service) CreateReservation(userID uint, req dto.CreateRequest) (*dto.Response, error) {

	reservation, err := s.repo.CreateReservationTransaction(
		userID,
		req.ZoneID,
		req.LicensePlate,
	)
	if err != nil {
		return nil, err
	}

	return &dto.Response{
		ID:             reservation.ID,
		UserID:         reservation.UserID,
		ZoneID:         reservation.ZoneID,
		LicensePlate:   reservation.LicensePlate,
		Status:         reservation.Status,
		CreatedAt:      reservation.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:      reservation.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
}

// ==========================================
// Get My Reservations
// ==========================================
func (s *service) GetMyReservations(userID uint) ([]dto.MyReservationResponse, error) {

	reservations, err := s.repo.GetMyReservations(userID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.MyReservationResponse, 0)

	for _, r := range reservations {

		response = append(response, dto.MyReservationResponse{
			ID:            r.ID,
			LicensePlate:  r.LicensePlate,
			Status:        r.Status,
			CreatedAt:     r.CreatedAt.Format("2006-01-02T15:04:05Z"),
			Zone: dto.ZoneResponse{
				ID:   r.Zone.ID,
				Name: r.Zone.Name,
				Type: r.Zone.Type,
			},
		})
	}

	return response, nil
}

// ==========================================
// Cancel Reservation
// ==========================================
func (s *service) CancelReservation(userID uint, reservationID uint) error {

	reservation, err := s.repo.GetByID(reservationID)
	if err != nil {
		return err
	}

	if reservation == nil {
		return errors.New("reservation not found")
	}

	if reservation.UserID != userID {
		return fmt.Errorf("forbidden")
	}

	if reservation.Status != StatusActive {
		return fmt.Errorf("reservation already %s", reservation.Status)
	}

	reservation.Status = StatusCancelled

	return s.repo.Update(reservation)
}

// ==========================================
// Admin - Get All Reservations
// ==========================================
func (s *service) GetAllReservations() ([]dto.AdminReservationResponse, error) {

	reservations, err := s.repo.GetAllReservations()
	if err != nil {
		return nil, err
	}

	response := make([]dto.AdminReservationResponse, 0)

	for _, r := range reservations {

		response = append(response, dto.AdminReservationResponse{
			ID:           r.ID,
			LicensePlate: r.LicensePlate,
			Status:       r.Status,
			CreatedAt:    r.CreatedAt.Format("2006-01-02T15:04:05Z"),

			User: dto.UserResponse{
				ID:    r.User.ID,
				Name:  r.User.Name,
				Email: r.User.Email,
				Role:  r.User.Role,
			},

			Zone: dto.ZoneResponse{
				ID:   r.Zone.ID,
				Name: r.Zone.Name,
				Type: r.Zone.Type,
			},
		})
	}

	return response, nil
}