package controllers

import (
	"net/http"

	"example.com/FanHubStore-api/initializers"
	"example.com/FanHubStore-api/models"
	"github.com/gin-gonic/gin"
)

func ProductCreate(c *gin.Context) {
	// get data off req body
	var product struct {
		ID    uint
		Name  string
		Price float64
		//  decimal.Decimal // Use decimal.Decimal from a library like "github.com/shopspring/decimal"
		Quantity    int
		Description string
		Category    string
	}

	// run c.Bind() and pass reference to that body
	c.Bind(&product)
	// create post
	post := models.CreateProductRequest{Name: product.Name, Price: product.Price, Quantity: product.Quantity, Description: product.Description, Category: product.Category}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"Post": product,
	})

}

func ProductIndex(c *gin.Context) {
	// get the product

	// create a variable to store our post
	var products []models.CreateProductRequest
	initializers.DB.Find(&products)
	// respond with them
	c.JSON(200, gin.H{
		"post": products,
	})

}

func ProductShow(c *gin.Context) {
	// Get the ID from the URL
	id := c.Param("id")

	// Create a variable to store the product
	var product models.CreateProductRequest // Assuming models.Product is your Product model struct

	// Retrieve the product from the database
	if err := initializers.DB.First(&product, id).Error; err != nil {
		// If the product is not found, return a 404 Not Found response
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Respond with the product
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func PostUpdate(c *gin.Context) {
	// Get the ID of the URL
	id := c.Param("id")

	// get data off request body
	var product struct {
		ID    uint
		Name  string
		Price float64
		//  decimal.Decimal // Use decimal.Decimal from a library like "github.com/shopspring/decimal"
		Quantity    int
		Description string
		Category    string
	}

	// run c.Bind() and pass reference to that body
	c.Bind(&product)
	// find the post were updating
	var post models.CreateProductRequest
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.CreateProductRequest{Name: product.Name, Price: product.Price, Quantity: product.Quantity, Description: product.Description, Category: product.Category})

	// respond with it
	c.JSON(http.StatusOK, gin.H{"product": product})
}

func ProductDelete(c *gin.Context) {
	// get id of url
	id := c.Param("id")

	// delete post
	initializers.DB.Delete(&models.CreateProductRequest{}, id)

	// respond with it
	c.Status(200)
}
