package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	_ "github.com/luis13005/pos-go/docs"
	"github.com/luis13005/pos-go/internal/dto"
	"github.com/luis13005/pos-go/internal/infra/database"
	"github.com/luis13005/pos-go/internal/model"
	"github.com/luis13005/pos-go/pkg/entidade"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(productDB database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: productDB}
}

// CreateProduct godoc
// @Sumary Create product
// @Description Create product
// @Tags product
// @Accept json
// @Produce json
// @Param request body dto.CreateProductDto true "product request"
// @Success 201
// @Failure 	500 	{object} 	error
// @Router /product [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductDto
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

// GetProductById godoc
// @Sumary 		get a product
// @Description get a product
// @Tags 		product
// @Accept 		json
// @Produce 	json
// @Param 		id 		path 		string 		true 	"product ID" Format(uuid)
// @Success 	200 	{object} 	model.Product
// @Failure 	404
// @Failure 	500 	{object} 	error
// @Router 		/product/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&product)
}

// GetAllProducts godoc
// @Sumary 		List products
// @Description get all products
// @Tags 		product
// @Accept 		json
// @Produce 	json
// @Param 		page 	query string false "page number"
// @Param 		limit 	query string false "limit"
// @Success 	200 	{array} model.Product
// @Failure 	404 	{object} 	error
// @Failure 	500 	{object} 	error
// @Router 		/product [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		pageNumber = 1
	}

	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		limitNumber = 0
	}

	products, err := h.ProductDB.FindAll(pageNumber, limitNumber, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// UpdateProduct godoc
// @Sumary 		Update a product
// @Description Update a product
// @Tags product
// @Accept 	json
// @Produce json
// @Param id 		path 	string 	true 		"product ID" 	Format(uuid)
// @Param request 	body 	dto.CreateProductDto true 	"product request"
// @Success 200
// @Failure 	404
// @Failure 	500 	{object} 	error
// @Router /product/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)

	_, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	product.ID, err = entidade.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DelteProduct godoc
// @Sumary 		Delete a product
// @Description Delete a product
// @Tags product
// @Accept 	json
// @Produce json
// @Param id 		path 	string 	true 		"product ID" 	Format(uuid)
// @Success 	200
// @Failure 	404
// @Failure 	500 	{object} 	error
// @Router /product/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DelteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
