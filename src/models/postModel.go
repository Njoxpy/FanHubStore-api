package models

import "gorm.io/gorm"

type CreateProductRequest struct {
	gorm.Model
	Name        string  
	Price       float64 
	Quantity    int     
	Description string  
	Category    string  
}
