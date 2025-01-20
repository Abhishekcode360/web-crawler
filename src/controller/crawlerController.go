package controller

import (
	service "Crawler/service"
	"Crawler/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrawlController(c *gin.Context) {
	var domains []string
	if err := c.BindJSON(&domains); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input format"})
	} else {
		result := service.CrawlDomains(c, domains)
		if err := utils.SaveToFile("output.json", result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save result to file"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "crawling completed, data saved to output.json"})
		}
	}
}
