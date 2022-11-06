package models

import "github.com/kingsbloc/auto-shop/internal/utils"

type Store struct {
	ID       string
	Products map[string]Product
}

func (s *Store) AddItem(item Item, price float64, quantity int64) {
	itemId := item.GetID()
	if p, e := s.Products[itemId]; e {
		p.Quantity = p.Quantity + quantity // increment quantity
		p.Price = price                    // update price if needed
		s.Products[itemId] = p
	} else {
		s.Products[itemId] = Product{
			ID:       utils.GenID(),
			Item:     item,
			Quantity: 1,
			Price:    price,
		}
	}

}
