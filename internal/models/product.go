package models

import "fmt"

type ProductItem interface {
	GetID() string
	GetName() string
	GetDetails() string
}

type Product struct {
	ID       string
	Price    float64
	Quantity int64
	Item     ProductItem
}

func (p *Product) Details() string {
	return fmt.Sprintf("Price: %f\nQuantity: %d\nItem:%q", p.Price, p.Quantity, p.Item)
}

func (p *Product) InStock() bool {
	return p.Quantity > 0
}
