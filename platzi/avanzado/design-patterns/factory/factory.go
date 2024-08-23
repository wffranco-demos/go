package factory

import (
	"errors"
	"fmt"
)

type Product interface {
	GetName() string
	SetName(name string)
	GetStock() int
	SetStock(stock int)
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) GetName() string {
	return c.name
}

func (c *Computer) SetName(name string) {
	c.name = name
}

func (c *Computer) GetStock() int {
	return c.stock
}

func (c *Computer) SetStock(stock int) {
	c.stock = stock
}

type Laptop struct {
	Computer
}

func NewLaptop(stock int) Product {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop",
			stock: stock,
		},
	}
}

type Desktop struct {
	Computer
}

func NewDesktop(stock int) Product {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop",
			stock: stock,
		},
	}
}

func GetComputerFactory(computerType string, stock int) (Product, error) {
	switch computerType {
	case "laptop":
		return NewLaptop(stock), nil
	case "desktop":
		return NewDesktop(stock), nil
	default:
		return nil, errors.New("invalid computer type")
	}
}

func PrintDetails(p Product) {
	fmt.Printf("Product: %s, Stock: %d\n", p.GetName(), p.GetStock())
}
