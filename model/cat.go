package model

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	ID    int
	Name  string `json:"name"`
	Breed string `json:"breed"`
}
