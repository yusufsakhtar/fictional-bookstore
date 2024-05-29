package models

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCompleted OrderStatus = "completed"
)

type Order struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	CartID string `json:"cart_id"`
	// TODO: need to represent items with their prices as well and separately, discounts
	// then, need to calculate total at order execution time.
	// That way, if items go OOS, we can still calculate the total based on the items that are available
	ItemIDs []string `json:"item_ids"`
	Total   float64  `json:"total"`
	// TODO: include logic for payment as part of order creation
	// PaymentMethodID string   `json:"payment_method_id"`

	// TODO: Shipping address eventually needs to be associated with order
	// ShippingAddressID string   `json:"shipping_address_id"`
	Status OrderStatus `json:"status"`
}
