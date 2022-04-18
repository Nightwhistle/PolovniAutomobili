package server

import (
	"stream-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	offers := new(controllers.OfferController)

	router.GET("/fetch", offers.Fetch)
	router.GET("/offer", offers.Index)

	return router
}
