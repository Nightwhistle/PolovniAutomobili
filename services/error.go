package services

import (
	"stream-api/models"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, code int, msg string) {
	error := models.CustomError{
		Code: code,
		Msg:  msg,
	}

	error.CreateErrorResponse(c)
}
