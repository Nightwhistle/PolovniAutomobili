package main

import (
	"stream-api/cache"
	"stream-api/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	cache.Init()
	server.Init()
}
