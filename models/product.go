package models

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