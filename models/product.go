package models

import (
	"fmt"
	"strings"
)


const minProductNameLen = 3

type Product struct {
	ID 	 string	 `bson:"_id,omitempty" json:"id"`
	Name string	 `bson:"name" json:"name"`
	SKU  string	 `bson:"sku" json:"sku"`
	Slug string	 `bson:"slug" json:"slug"`
}

type CreateProductRequest struct {
	SKU string `json:"sku"`
	Name string	`json:"name"`
}

//function to create new product from the request
func NewProductFromRequest(req *CreateProductRequest) (*Product, error){
	if err := validateCreateProductRequest(req); err != nil {
		return nil, err
	}

	parts := strings.Split(strings.ToLower(req.Name), " ")
	slug := strings.Join(parts, "-")

	return &Product{
		SKU: req.SKU,
		Name: req.Name,
		Slug: slug,
	}, nil
}

// function to validate the product to create
func validateCreateProductRequest(req *CreateProductRequest) error {
	if len(req.SKU) < 3 {
		return fmt.Errorf("The SKU of the product is short")
	}
	if len(req.Name) < minProductNameLen{
		return fmt.Errorf("The name of the product is too short")
	}
	return nil
}