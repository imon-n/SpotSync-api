package main

import (
	"goproject/internal/config"
	"goproject/internal/server"
)

func main() {
	cfg := config.LoadEnv()
	db := config.ConnectDatabase(cfg)
	server.Start(db, cfg)
}