package server

import (
	"goproject/internal/config"
	"goproject/internal/domain/parkingzone"
	"goproject/internal/domain/reservation"
	"goproject/internal/domain/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Start(db *gorm.DB, cfg *config.Config) {
	db.AutoMigrate(	&user.User{},
		&parkingzone.ParkingZone{},
		&reservation.Reservation{},)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(200, map[string]any{
			"success": true,
			"message": "SpotSync API is running successfully 🚀",
		})
	})

	user.RegisterRoutes(e, db, cfg)
	parkingzone.RegisterRoutes(e, db, cfg)
	reservation.RegisterRoutes(e, db, cfg)
	e.Start(":" + cfg.Port)
}
