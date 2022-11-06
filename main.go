package main

import (
	"github.com/kingsbloc/auto-shop/internal/services"
)

func main() {
	shopService := services.NewShopService()
	car1 := shopService.AddCar("BMW", "2 Series", 2020, "White")
	car2 := shopService.AddCar("Toyota", "Camry", 2021, "Red")
	car3 := shopService.AddCar("Tesla", "Model 3", 2019, "White")
	car4 := shopService.AddCar("Mercedes", "AMG SL", 2022, "Gray")
	car5 := shopService.AddCar("Lamborghini", "Urus", 2020, "Yellow")
	shopService.AddProduct(car1, 20000, 10)
	shopService.AddProduct(car2, 30000, 20)
	shopService.AddProduct(car3, 40000, 30)
	shopService.AddProduct(car4, 50000, 40)
	shopService.AddProduct(car5, 60000, 50)
}
