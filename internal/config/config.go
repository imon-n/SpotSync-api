// package config

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// type Config struct {
// 	Port      string
// 	Dsn       string
// 	JwtSecret string
// }

// func LoadEnv() *Config {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	return &Config{
// 		Port:      os.Getenv("PORT"),
// 		Dsn:       os.Getenv("DSN"),
// 		JwtSecret: os.Getenv("JWT_SECRET"),
// 	}
// }


package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	Dsn       string
	JwtSecret string
}

func LoadEnv() *Config {

	// Local development-এর জন্য .env load করবে।
	// Render-এ .env না থাকলেও app চলবে।
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using Render environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:      port,
		Dsn:       os.Getenv("DSN"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}