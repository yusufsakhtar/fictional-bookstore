package models

type PaymentMethodType string

const (
	PaymentMethodTypeCreditCard PaymentMethodType = "credit_card"
	PaymentMethodTypeBalance    PaymentMethodType = "balance"
)

type PaymentMethod struct {
	ID      string            `json:"id"`
	Type    PaymentMethodType `json:"type"`
	Balance float64           `json:"balance"`
}

func (p *PaymentMethod) IsValid() bool {
	if p.Type == PaymentMethodTypeBalance && p.Balance < 0 {
		return false
	}
	return true
}
