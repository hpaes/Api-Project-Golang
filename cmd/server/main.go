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
	"github.com/go-chi/jwtauth/v5"
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
	userHandler := handlers.NewUserHandler(userDb)

	router := chi.NewRouter()
	router.Use(middleware.Heartbeat("/"))
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(300 * time.Millisecond))
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwt", config.TokenAuth))
	router.Use(middleware.WithValue("JwtExpiresIn", config.JwtExpiresIn))

	router.Route("/users", func(router chi.Router) {
		router.Post("/", userHandler.CreateUser)
		router.Post("/generate_token", userHandler.GetJwtInput)

	})

	router.Route("/products", func(router chi.Router) {
		router.Use(jwtauth.Verifier(config.TokenAuth))
		router.Use(jwtauth.Authenticator)
		router.Get("/", productHandler.GetProducts)
		router.Post("/", productHandler.CreateProduct)
		router.Get("/{id}", productHandler.GetProduct)
		router.Put("/{id}", productHandler.UpdateProduct)
		router.Delete("/{id}", productHandler.DeleteProduct)
	})

	http.ListenAndServe(":8080", router)
}
