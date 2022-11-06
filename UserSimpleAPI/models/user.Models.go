package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
