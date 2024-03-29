package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetNews(c *gin.Context) {
	newsApiKey := os.Getenv("NEWS_API_KEY")
	category := c.Param("category")
	resp, err := http.Get("https://newsapi.org/v2/everything?q=" + category + "&apiKey=" + newsApiKey)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(resp.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.Data(200, "application/json", respbody)
}
