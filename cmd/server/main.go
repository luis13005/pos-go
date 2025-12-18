package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/luis13005/pos-go/configs"
	"github.com/luis13005/pos-go/internal/infra/database"
	"github.com/luis13005/pos-go/internal/infra/webserver/handlers"
	"github.com/luis13005/pos-go/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs := configs.LoadConfig(".")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{}, &model.Product{})
	fmt.Println(configs)
	productDb := database.NewProductDB(db)
	userDB := database.NewUserDB(db)
	productHandler := handlers.NewProductHandler(productDb)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/product", productHandler.CreateProductHandler)
	r.Get("/product/{id}", productHandler.GetProductById)
	r.Get("/product", productHandler.GetAllProducts)
	r.Put("/product/{id}", productHandler.UpdateProduct)
	r.Delete("/product/{id}", productHandler.DelteProduct)

	r.Post("/user", userHandler.CreateUser)
	http.ListenAndServe(":8000", r)
}
