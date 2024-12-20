package product

import "fmt"

type Product struct {
	Code     string
	Name     string
	Price    float64
	Quantity int
}

func NewProduct(code, name string, price float64, quantity int) *Product {
	return &Product{
		Code:     code,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

func (p Product) String() string {
	return fmt.Sprintf("| %-8s | %-25s | %12.0f | %10d |\n", p.Code, p.Name, p.Price, p.Quantity)
}
