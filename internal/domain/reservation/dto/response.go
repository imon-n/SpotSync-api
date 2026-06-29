package dto

// ================================
// Create Reservation Response
// ================================

type Response struct {
	ID uint `json:"id"`

	UserID uint `json:"user_id,omitempty"`

	ZoneID uint `json:"zone_id,omitempty"`

	LicensePlate string `json:"license_plate"`

	Status string `json:"status"`

	CreatedAt string `json:"created_at"`

	UpdatedAt string `json:"updated_at,omitempty"`
}

// ================================
// Zone Response
// ================================

type ZoneResponse struct {
	ID uint `json:"id"`

	Name string `json:"name"`

	Type string `json:"type"`
}

// ================================
// User Response
// ================================

type UserResponse struct {
	ID uint `json:"id"`

	Name string `json:"name"`

	Email string `json:"email"`

	Role string `json:"role"`
}

// ================================
// GET /my-reservations
// ================================

type MyReservationResponse struct {
	ID uint `json:"id"`

	LicensePlate string `json:"license_plate"`

	Status string `json:"status"`

	Zone ZoneResponse `json:"zone"`

	CreatedAt string `json:"created_at"`
}

// ================================
// GET /reservations (Admin)
// ================================

type AdminReservationResponse struct {
	ID uint `json:"id"`

	LicensePlate string `json:"license_plate"`

	Status string `json:"status"`

	User UserResponse `json:"user"`

	Zone ZoneResponse `json:"zone"`

	CreatedAt string `json:"created_at"`
}