package parkingzone

import (
	"net/http"
	"strconv"

	"goproject/internal/domain/parkingzone/dto"
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

// POST /api/v1/zones
func (h *handler) CreateZone(c *echo.Context) error {

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

	res, err := h.service.CreateZone(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"success": true,
		"message": "Parking zone created successfully",
		"data":    res,
	})
}

// GET /api/v1/zones
func (h *handler) GetAllZones(c *echo.Context) error {

	res, err := h.service.GetAllZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"message": "Parking zones retrieved successfully",
		"data":    res,
	})
}

// GET /api/v1/zones/:id
func (h *handler) GetZoneByID(c *echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid parking zone ID",
		})
	}

	res, err := h.service.GetZoneByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"message": "Parking zone retrieved successfully",
		"data":    res,
	})
}