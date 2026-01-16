package api

import (
	"encoding/json"
	"net/http"

	"github.com/jimmymuthoni/go_ecomm_backend/models"
	"github.com/jimmymuthoni/go_ecomm_backend/store"
	"github.com/jimmymuthoni/weavebox"
)

type ProductHandler struct {
	store 	store.ProductStorer
}

func NewProductHandler(pStore store.ProductStorer) *ProductHandler {
	return &ProductHandler{
		store: pStore,
	}
}

func (h *ProductHandler) HandlePostProduct (c *weaviate.Context) error{
	productReq := &models.CreateProductRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(productReq); err != nil {
		return err
	}
	product, err := models.NewProductFromRequest(productReq)
	if err != nil {
		return err
	}

	if err := h.store.Insert(c.Context, product); err != nil {
		return err
	}
	return c.Json(http.StatusOK, product)
}

func (h *ProductHandler) HandleGetProducts(c *weavebox.Context) error{
	products, err := h.store.GetAll(c.Context)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}

