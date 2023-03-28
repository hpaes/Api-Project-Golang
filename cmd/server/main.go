package main

import (
	"log"

	"github.com/hpaes/api-project-golang/configs"
	"github.com/hpaes/api-project-golang/internal/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config := configs.LoadConfig()
	println(config.DBDriver)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})
}
