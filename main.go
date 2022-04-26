package main

import (
	"stream-api/cache"
	"stream-api/server"
	"stream-api/services"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	services.OpenBrowser("http://localhost:8080/fetch")

	cache.Init()
	server.Init()
}
