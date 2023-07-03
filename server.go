package main

import (
	"github.com/jinzhu/gorm"
)

type Server struct {
	gorm.Model // extends gorm model (updated/deleted_at etc)
	Name   string
	Type   string
	Status string
}
