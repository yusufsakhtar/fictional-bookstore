package models

type Cart struct {
	UserID  string   `json:"user_id"`
	ID      string   `json:"id"`
	ItemIds []string `json:"item_ids"`
}
