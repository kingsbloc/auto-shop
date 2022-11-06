package models

import (
	"errors"
	"time"

	"github.com/kingsbloc/auto-shop/internal/utils"
)

type Store struct {
	ID       string
	Products map[string]Product
	Orders   []Order
}

func (s *Store) GetItem(itemId string) *Product {
	if p, e := s.Products[itemId]; e {
		return &p
	}
	return nil
}

func (s *Store) AddItem(item ProductItem, price float64, quantity int64) {
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

// Listing all product items in the store
func (s *Store) ListItems(inStockOnly bool) []ProductItem {
	var items []ProductItem
	for _, v := range s.Products {
		if inStockOnly {
			if v.InStock() {
				items = append(items, v.Item)
			}
		} else {
			items = append(items, v.Item)
		}
	}
	return items
}

// Sell an item
func (s *Store) SellItem(itemId string, quantity int64) (bool, error) {
	p := s.GetItem(itemId)
	if p != nil {
		if !p.InStock() {
			return false, errors.New("item not in stock")
		}
		if quantity > p.Quantity {
			return false, errors.New("insufficent items")
		}
		p.Quantity = p.Quantity - quantity
		s.Products[itemId] = *p
		s.Orders = append(s.Orders, Order{p, p.Price, time.Now()})
		return true, nil
	}
	return false, errors.New("item not found")
}

// Show a list of sold items and the total price
func (s *Store) ListOrders() ([]Order, float64) {
	total := float64(0)
	for _, v := range s.Orders {
		total += v.SalePrice
	}
	return s.Orders, total
}
