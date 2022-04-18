package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := NewRouter()
	r.Run(":" + os.Getenv("SERVER_PORT"))
}

func CreateError(msg string) interface{} {
	return gin.H{
		"error": msg,
	}
}
