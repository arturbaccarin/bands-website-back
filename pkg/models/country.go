package models

import "gorm.io/gorm"

// Country contains the fields of the table
type Country struct {
	gorm.Model
	ID        uint
	Name      string
	Continent string
}
