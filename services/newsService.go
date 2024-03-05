package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
)

func GetNews(c *gin.Context) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
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
	fmt.Println("reponse", respbody)
	c.Data(200, "application/json", respbody)
}
