package store

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/KhoaDoesTech/ProductManagement/product"
)

type Store struct {
	Products []product.Product
}

type StoreManager interface {
	Add(product.Product) error
	Display()
	Search(name string) (*product.Product, error)
	Delete(code string) error
	TotalValue() float64
}

// Design Pattern: Singleton

var instance *Store
var once sync.Once

func GetStoreInstance() *Store {
	once.Do(func() {
		instance = &Store{
			Products: make([]product.Product, 0),
		}
	})
	return instance
}

func (s *Store) Add(p product.Product) error {
	if strings.TrimSpace(p.Code) == "" || strings.TrimSpace(p.Name) == "" {
		return errors.New("mã sản phẩm và tên sản phẩm không được để trống")
	}

	for _, prod := range s.Products {
		if prod.Code == p.Code {
			return errors.New("sản phẩm đã tồn tại với mã này")
		}
		if strings.EqualFold(prod.Name, p.Name) {
			return errors.New("sản phẩm đã tồn tại với tên này")
		}
	}

	s.Products = append(s.Products, p)

	return nil
}

func (s *Store) Display() {
	if s.IsEmpty() {
		fmt.Println("\nCửa hàng hiện không có sản phẩm nào.")
		return
	}

	printProducts(s.Products)
}

func DisplayProducts(products []product.Product) {
	if len(products) == 0 {
		fmt.Println("Danh sách sản phẩm trống.")
		return
	}

	printProducts(products)
}

func printProducts(products []product.Product) {
	fmt.Println("\nDanh sách sản phẩm:")
	fmt.Println("+----------+---------------------------+--------------+------------+")
	fmt.Println("| MÃ SP    | TÊN SẢN PHẨM              | GIÁ (VND)    | SỐ LƯỢNG   |")
	fmt.Println("+----------+---------------------------+--------------+------------+")
	for _, prod := range products {
		fmt.Print(prod)
	}
	fmt.Println("+----------+---------------------------+--------------+------------+")
}

func (s *Store) Search(keyword string) ([]product.Product, error) {
	if s.IsEmpty() {
		return nil, errors.New("không có sản phẩm nào trong cửa hàng để tìm kiếm")
	}

	keyword = strings.ToLower(strings.TrimSpace(keyword))
	if keyword == "" {
		return nil, errors.New("từ khóa tìm kiếm không được để trống")
	}

	var result []product.Product
	for _, prod := range s.Products {
		if strings.Contains(strings.ToLower(prod.Name), keyword) {
			result = append(result, prod)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("không tìm thấy sản phẩm với từ khóa này")
	}

	return result, nil
}

func (s *Store) Delete(code string) error {
	if s.IsEmpty() {
		return errors.New("không có sản phẩm nào trong cửa hàng để xóa")
	}

	code = strings.TrimSpace(code)
	if code == "" {
		return errors.New("mã sản phẩm không được để trống")
	}

	for i, prod := range s.Products {
		if prod.Code == code {
			s.Products = append(s.Products[:i], s.Products[i+1:]...) // Lấy từ 0 -> 2 và lấy từ 4 -> cuối => Xóa 3
			return nil
		}
	}

	return errors.New("không tìm thấy sản phẩm với mã này")
}

func (s *Store) TotalValue() float64 {
	if s.IsEmpty() {
		fmt.Println("\nKhông có sản phẩm nào để tính giá trị.")
		return 0
	}

	var total float64
	for _, prod := range s.Products {
		total += prod.Price * float64(prod.Quantity)
	}

	return total
}

func (s *Store) IsEmpty() bool {
	return len(s.Products) == 0
}
