package main

import (
	"github.com/jinzhu/gorm"
)

type Server struct {
	gorm.Model // extends gorm model (updated/deleted_at etc)
	ID     string `gorm:"primary_key"` // decorator to tell gorm to use this as primary key
	Name   string
	Type   string
	Status string
}
