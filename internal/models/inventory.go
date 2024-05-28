package models

const (
	InventoryItemStatusAvailable   = "available"
	InventoryItemStatusPendingSale = "pending_sale"
)

type InventoryItem struct {
	Item  *Item               `json:"item"`
	Stock *InventoryItemStock `json:"stock"`
}

type InventoryItemStock struct {
	// We could potentially further separate Available into in_cart and fully_available
	Available   int `json:"available"`
	PendingSale int `json:"pending_sale"`
}
