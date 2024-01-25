package main

import (
	"fmt"

	"example.com/FanHubStore-api/initializers"
	"github.com/gin-gonic/gin"
)

// function for loading environment variables
func init() {
	initializers.LoadEnvVariables()
}

func main() {

	// load environment variables

	fmt.Println("Hello")

	/* 
	1. r := gin.Default() Creates a route
	2. inner code adds a route
	3. r.Run() runs the route of the path
	*/

	r := gin.Default() 
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
