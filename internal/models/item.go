package models

type Item struct {
	DisplayName string  `json:"display_name"`
	SKU         string  `json:"sku"`
	Price       float64 `json:"price"`
}
