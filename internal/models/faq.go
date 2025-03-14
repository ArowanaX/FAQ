package models

import "gorm.io/gorm"

type FaqModels struct {
	gorm.Model
	Category Category
	Question string `gorm:"size:255;not null"`
	Answer   string `gorm:"size:255;not null"`
}

type Category struct {
	Title string `gorm:"size:20;not null"`
}
