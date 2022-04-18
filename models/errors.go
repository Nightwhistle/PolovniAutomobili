package models

import "github.com/gin-gonic/gin"

type CustomError struct {
	Code int
	Msg  string
}

func (e *CustomError) CreateErrorResponse(c *gin.Context) {
	c.JSON(e.Code, gin.H{
		"error": e.Msg,
	})
}
