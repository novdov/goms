package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	MapUrls()

	log.Fatal(router.Run(":8080"))
}
