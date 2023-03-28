package main

import (
	"log"
	"net/http"

	"github.com/hpaes/api-project-golang/configs"
	"github.com/hpaes/api-project-golang/internal/entity"
	"github.com/hpaes/api-project-golang/internal/infra/database"
	"github.com/hpaes/api-project-golang/internal/infra/webservers/handlers"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs.LoadConfig()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDb := database.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productDb)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8080", nil)
}
