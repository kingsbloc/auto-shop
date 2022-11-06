package models

import "fmt"

type Item interface {
	GetID() string
	Name() string
	GetDetails() string
}

type Product struct {
	ID       string
	Price    float64
	Quantity int64
	Item     Item
}

func (p *Product) Details() string {
	return fmt.Sprintf("Price: %f\nQuantity: %d\nItem:%q", p.Price, p.Quantity, p.Item)
}

func (p *Product) InStock() bool {
	return p.Quantity > 0
}
