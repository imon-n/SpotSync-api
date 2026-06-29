package reservation

import (
	"net/http"
	"strconv"

	"goproject/internal/domain/reservation/dto"
	"goproject/internal/httpresponse"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(s *service) *handler {
	return &handler{
		service: s,
	}
}

// POST /api/v1/reservations
func (h *handler) CreateReservation(c *echo.Context) error {

	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	var req dto.CreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	res, err := h.service.CreateReservation(userID, req)
	if err != nil {
		return c.JSON(http.StatusConflict, httpresponse.Error{
			Code:    http.StatusConflict,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"success": true,
		"message": "Reservation confirmed successfully",
		"data":    res,
	})
}

// GET /api/v1/reservations/my-reservations
func (h *handler) GetMyReservations(c *echo.Context) error {

	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	res, err := h.service.GetMyReservations(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"message": "My reservations retrieved successfully",
		"data":    res,
	})
}

// DELETE /api/v1/reservations/:id
func (h *handler) CancelReservation(c *echo.Context) error {

	userID, ok := c.Get("user_id").(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.Error{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid reservation ID",
		})
	}

	err = h.service.CancelReservation(userID, uint(id))
	if err != nil {

		if err.Error() == "forbidden" {
			return c.JSON(http.StatusForbidden, httpresponse.Error{
				Code:    http.StatusForbidden,
				Message: "You can cancel only your own reservation",
			})
		}

		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"message": "Reservation cancelled successfully",
	})
}

// GET /api/v1/reservations
func (h *handler) GetAllReservations(c *echo.Context) error {

	res, err := h.service.GetAllReservations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"message": "All reservations retrieved successfully",
		"data":    res,
	})
}