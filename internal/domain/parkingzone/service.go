package parkingzone

import (
	"fmt"

	"goproject/internal/domain/parkingzone/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

// Create Parking Zone
func (s *service) CreateZone(req dto.CreateRequest) (*dto.Response, error) {

	if req.Type != "general" &&
		req.Type != "ev_charging" &&
		req.Type != "covered" {
		return nil, fmt.Errorf("invalid parking zone type")
	}

	zone := ParkingZone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	if err := s.repo.Create(&zone); err != nil {
		return nil, err
	}

	return &dto.Response{
		ID:             zone.ID,
		Name:           zone.Name,
		Type:           zone.Type,
		TotalCapacity:  zone.TotalCapacity,
		AvailableSpots: zone.TotalCapacity,
		PricePerHour:   zone.PricePerHour,
		CreatedAt:      zone.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:      zone.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

// Get All Parking Zones
func (s *service) GetAllZones() ([]dto.Response, error) {

	zones, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dto.Response, 0, len(zones))

	for _, zone := range zones {

		responses = append(responses, dto.Response{
			ID:             zone.ID,
			Name:           zone.Name,
			Type:           zone.Type,
			TotalCapacity:  zone.TotalCapacity,
			AvailableSpots: zone.TotalCapacity, // Update after Reservation module
			PricePerHour:   zone.PricePerHour,
			CreatedAt:      zone.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return responses, nil
}

// Get Single Parking Zone
func (s *service) GetZoneByID(id uint) (*dto.Response, error) {

	zone, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if zone == nil {
		return nil, fmt.Errorf("parking zone not found")
	}

	return &dto.Response{
		ID:             zone.ID,
		Name:           zone.Name,
		Type:           zone.Type,
		TotalCapacity:  zone.TotalCapacity,
		AvailableSpots: zone.TotalCapacity, // Update after Reservation module
		PricePerHour:   zone.PricePerHour,
		CreatedAt:      zone.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:      zone.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}