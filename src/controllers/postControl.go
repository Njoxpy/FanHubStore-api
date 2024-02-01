package controllers

import (
	"example.com/FanHubStore-api/initializers"
	"example.com/FanHubStore-api/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data off req body

	// create post
	post := models.CreateProductRequest{Name: "Manchester United Home Jersey 2022/23", Price:  79.99, Quantity: 100, Description: "Official home jersey of Manchester United for the 2022/23 season.", Category: "Sports Apparel"}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	
	// return it
	c.JSON(200, gin.H{
		"Hello Njox": post,
	})

}
