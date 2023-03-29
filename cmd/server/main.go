package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hpaes/api-project-golang/configs"
	"github.com/hpaes/api-project-golang/internal/entity"
	"github.com/hpaes/api-project-golang/internal/infra/database"
	"github.com/hpaes/api-project-golang/internal/infra/webservers/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config := configs.LoadConfig()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDb := database.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productDb)

	userDb := database.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userDb, config.TokenAuth, config.JwtExpiresIn)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(300 * time.Millisecond))

	router.Route("/products", func(r chi.Router) {
		router.Post("/", productHandler.CreateProduct)
		router.Get("/{id}", productHandler.GetProduct)
		router.Get("/", productHandler.GetProducts)
		router.Put("/{id}", productHandler.UpdateProduct)
		router.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Route("/users", func(r chi.Router) {
		router.Post("/users", userHandler.CreateUser)
		router.Post("/users/generate_token", userHandler.GetJwtInput)

	})

	http.ListenAndServe(":8080", router)
}
