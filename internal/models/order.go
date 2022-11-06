package models

import "time"

type Order struct {
	Product   *Product
	SalePrice float64
	Date      time.Time
}
