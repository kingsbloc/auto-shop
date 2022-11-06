package services

import (
	"fmt"

	"github.com/kingsbloc/auto-shop/internal/models"
	"github.com/kingsbloc/auto-shop/internal/utils"
)

type ShopService interface {
	AddCar(brand string, model string, year int64, color string) *models.Car
	AddProduct(item models.Item, price float64, quantity int64)
}

type shopService struct {
	store *models.Store
}

func NewShopService() ShopService {
	store := models.Store{
		ID:       utils.GenID(),
		Products: make(map[string]models.Product),
	}
	return &shopService{store: &store}
}

func (s *shopService) AddCar(brand string, model string, year int64, color string) *models.Car {
	car := &models.Car{Brand: brand, Model: model, Year: year, Color: color, ID: utils.GenID()}
	return car
}

func (s *shopService) AddProduct(item models.Item, price float64, quantity int64) {
	s.store.AddItem(item, price, quantity)
	fmt.Println("Product added to Store")
}
