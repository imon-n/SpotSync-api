package reservation

import (
	"goproject/internal/auth"
	"goproject/internal/config"
	"goproject/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {

	// Dependencies
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	// JWT
	jwtService := auth.NewJWTService(cfg.JwtSecret)

	// Protected Routes
	api := e.Group("/api/v1/reservations")
	api.Use(middlewares.AuthMiddleware(jwtService))

	// Driver + Admin
	api.POST("", handler.CreateReservation)
	api.GET("/my-reservations", handler.GetMyReservations)
	api.DELETE("/:id", handler.CancelReservation)

	// Admin Only
	admin := api.Group("")
	admin.Use(middlewares.AdminMiddleware())

	admin.GET("", handler.GetAllReservations)
}