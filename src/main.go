package main

import (
	"example.com/FanHubStore-api/controllers"
	"example.com/FanHubStore-api/initializers"
	"github.com/gin-gonic/gin"
)

// function for loading environment variables
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	// return it
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.Run()
}
