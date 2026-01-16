package store

import (
	"context"

	"github.com/jimmymuthoni/go_ecomm_backend/models"
)

type ProductStorer interface {
	Insert(context.Context, *models.Product) error
	GetAll(context.Context)([]*models.Product, error)
	GetByID(context.Context, string)(*models.Product, error)
}