package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Post struct {
	gorm.Model
	Title string `gorm:"unique" json:"title"`
	Text  string `json:"text"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.LogMode(true)
	db.AutoMigrate(&Post{})
	return db
}
