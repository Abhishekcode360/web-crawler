package main

import (
	"fmt"

	controller "Crawler/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting the Server......")
	router := gin.New()

	router.POST("/webcrawl", controller.CrawlController)

	router.Run(":8080")
}
