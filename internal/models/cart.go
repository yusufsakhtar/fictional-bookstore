package models

type Cart struct {
	Items []Item `json:"items"`
}

func (c *Cart) AddItem(i Item) {
	c.Items = append(c.Items, i)
}
