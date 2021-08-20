package database

import (
	"gorm.io/gorm"
)

type DBModel struct {
	DB *gorm.DB
}
