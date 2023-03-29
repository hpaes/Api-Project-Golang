package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hpaes/api-project-golang/internal/entity"
	"github.com/hpaes/api-project-golang/internal/infra/database"
	"github.com/hpaes/api-project-golang/internal/infra/webservers/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// configs.LoadConfig()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDb := database.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productDb)

	userDb := database.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userDb)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(300 * time.Millisecond))

	router.Post("/products", productHandler.CreateProduct)
	router.Get("/products/{id}", productHandler.GetProduct)
	router.Get("/products", productHandler.GetProducts)
	router.Put("/products/{id}", productHandler.UpdateProduct)
	router.Delete("/products/{id}", productHandler.DeleteProduct)

	router.Post("/users", userHandler.CreateUser)

	http.ListenAndServe(":8080", router)
}
