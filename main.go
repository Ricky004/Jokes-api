package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", getValues)
	r.Run(":9090")

}

var url = "https://v2.jokeapi.dev/joke/Any"

func getValues(c *gin.Context) {
	res, err := http.Get(url)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	var target map[string]interface{}
	err = json.Unmarshal(data, &target)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
	}

	
	c.JSON(http.StatusOK, gin.H{
		"setup": target["setup"],
		"delivery": target["delivery"],
	}) // Send the actual joke data

}
