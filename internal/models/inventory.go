package models

type Inventory struct {
	Items []Item `json:"items"`
}

func (i *Inventory) AddItem(item Item) {
	i.Items = append(i.Items, item)
}

func (i *Inventory) GetItems() []Item {
	return i.Items
}

func (i *Inventory) GetItemBySKU(sku string) Item {
	for _, i := range i.Items {
		if i.SKU == sku {
			return i
		}
	}
	return Item{}
}
