package parkingzone

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

	// JWT Service
	jwtService := auth.NewJWTService(cfg.JwtSecret)

	// Public Routes
	e.GET("/api/v1/zones", handler.GetAllZones)
	e.GET("/api/v1/zones/:id", handler.GetZoneByID)

	// Admin Routes
	admin := e.Group("/api/v1/zones")
	admin.Use(middlewares.AuthMiddleware(jwtService))
	admin.Use(middlewares.AdminMiddleware())

	admin.POST("", handler.CreateZone)
}