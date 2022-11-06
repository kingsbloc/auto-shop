package services

import (
	"github.com/kingsbloc/auto-shop/internal/models"
	"github.com/kingsbloc/auto-shop/internal/utils"
)

type ShopService interface {
	AddCar(brand string, model string, year int64, color string) *models.Car
	AddProduct(item models.ProductItem, price float64, quantity int64)
	ListItems(inStockOnly bool) []models.ProductItem
	ListOrders() ([]models.Order, float64)
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

// number of cars he has sold and Sum total of the prices of cars he has sold
func (s *shopService) ListOrders() ([]models.Order, float64) {
	total := float64(0)
	for _, v := range s.store.Orders {
		total += v.SalePrice
	}
	return s.store.Orders, total
}

func (s *shopService) ListItems(inStockOnly bool) []models.ProductItem {
	return s.store.ListItems(inStockOnly)
}

// Add Product
func (s *shopService) AddProduct(item models.ProductItem, price float64, quantity int64) {
	s.store.AddItem(item, price, quantity)
}

// number of cars that are left to be sold and sum of the prices of the cars left
func (s *shopService) GetUnsold() (int, float64) {
	var products []models.ProductItem
	total := float64(0)

	for _, v := range s.store.Products {
		if v.InStock() {
			total += v.Price * float64(v.Quantity)
			products = append(products, v.Item)
		}
	}
	return len(products), total
}
