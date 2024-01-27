package models

import "gorm.io/gorm"

type CreateProductRequest struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	Description string  `json:"description"`
}

/*
Steps by step

1. See How to declare models in go
2.

*/
