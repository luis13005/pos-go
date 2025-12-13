package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luis13005/pos-go/configs"
	"github.com/luis13005/pos-go/internal/dto"
	"github.com/luis13005/pos-go/internal/infra/database"
	"github.com/luis13005/pos-go/internal/model"
	"github.com/shopspring/decimal"
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
	productHandler := NewProductHandler(productDb)

	http.HandleFunc("/product", productHandler.CreateProductHandler)
	http.ListenAndServe(":8000", nil)
}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(productDB database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: productDB}
}

func (h ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product dto.ProductDto
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	newProdudct, err := model.NewProduct(&model.Product{Nome: product.Nome, Preco: decimal.NewFromFloat32(product.Preco)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = h.ProductDB.CreateProduct(newProdudct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
