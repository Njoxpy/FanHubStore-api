package main

import (
	"example.com/FanHubStore-api/initializers"
	"example.com/FanHubStore-api/models"
)

// look on how to migrate database

func init() {
	// connect to the database and load Env variables
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {

	/* db.AutoMigrate stands for database but inorder to access the database then use intializers.db and then pass structs over our models
	 */
	initializers.DB.AutoMigrate(&models.CreateProductRequest{})
}
