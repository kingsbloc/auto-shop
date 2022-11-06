package models

import "fmt"

type Car struct {
	ID    string
	Brand string
	Year  int64
	Model string
	Color string
}

func (c *Car) GetName() string {
	return fmt.Sprintf("%s - %s - %d", c.Brand, c.Model, c.Year)
}

func (c *Car) GetID() string {
	return c.ID
}

func (c *Car) GetDetails() string {
	return fmt.Sprintf("ID: %s\nBrand: %s\n Model: %s\nYear: %d\nColor: %s\n", c.ID, c.Brand, c.Model, c.Year, c.Color)
}
