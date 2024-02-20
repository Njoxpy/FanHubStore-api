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
	r.POST("/posts", controllers.ProductCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts", controllers.ProductIndex)
	r.GET("/posts/:id", controllers.ProductShow)
	r.DELETE("/posts/:id", controllers.ProductDelete)
	r.Run()
}
