package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Title string  `json:"name" binding:"required" gorm:"size:255;not null;"`
	Email string  `json:"email" binding:"required" gorm:"size:255;not null;"`
	Price float64 `json:"price" binding:"required" gorm:"size:255;not null;"`
}
