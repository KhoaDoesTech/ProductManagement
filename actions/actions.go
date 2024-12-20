package actions

import (
	"fmt"
	"os"
	"strings"

	"github.com/KhoaDoesTech/ProductManagement/product"
	"github.com/KhoaDoesTech/ProductManagement/store"
)

// Design Pattern: Strategy

type Action interface {
	Execute()
}

type AddProduct struct{}

func (a AddProduct) Execute() {
	var code, name string
	var price float64
	var quantity int

	fmt.Print("Nhập mã sản phẩm: ")
	fmt.Scan(&code)

	fmt.Print("Nhập tên sản phẩm: ")
	fmt.Scan(&name)

	name = strings.TrimSpace(name)

	for {
		fmt.Print("Nhập giá sản phẩm: ")
		fmt.Scan(&price)

		if price < 0 {
			fmt.Println("Lỗi::: Giá sản phẩm không được nhỏ hơn 0. Vui lòng nhập lại.")
		} else {
			break
		}
	}

	for {
		fmt.Print("Nhập số lượng sản phẩm: ")
		fmt.Scan(&quantity)

		if quantity < 0 {
			fmt.Println("Lỗi::: Số lượng sản phẩm phải lớn hơn 0. Vui lòng nhập lại.")
		} else {
			break
		}
	}

	prod := product.NewProduct(code, name, price, quantity)
	err := store.GetStoreInstance().Add(*prod)

	if err != nil {
		fmt.Println("Lỗi:::", err)
	} else {
		fmt.Println("Thêm sản phẩm thành công!")
	}
}

type DisplayProducts struct{}

func (d DisplayProducts) Execute() {
	store.GetStoreInstance().Display()
}

type SearchProduct struct{}

func (s SearchProduct) Execute() {
	var name string
	fmt.Print("Nhập tên sản phẩm cần tìm: ")
	fmt.Scan(&name)

	foundProducts, err := store.GetStoreInstance().Search(name)
	if err != nil {
		fmt.Println("Lỗi:::", err)
	} else {
		fmt.Println("Tìm thấy các sản phẩm sau:")
		store.DisplayProducts(foundProducts)
	}
}

type DeleteProduct struct{}

func (s DeleteProduct) Execute() {
	var code string

	fmt.Print("Nhập mã sản phẩm cần xóa: ")
	fmt.Scan(&code)

	err := store.GetStoreInstance().Delete(code)
	if err != nil {
		fmt.Println("Lỗi:", err)
	} else {
		fmt.Println("Xóa sản phẩm thành công!")
	}
}

type CalculateTotalValue struct{}

func (c CalculateTotalValue) Execute() {
	total := store.GetStoreInstance().TotalValue()
	fmt.Printf("Tổng giá trị sản phẩm trong cửa hàng: %.0f VND\n", total)
}

type ExitApplication struct{}

func (e ExitApplication) Execute() {
	fmt.Println("Đã thoát chương trình.")
	os.Exit(0)
}

type InitData struct{}

func (i InitData) Execute() {
	storeInstance := store.GetStoreInstance()

	sampleProducts := []product.Product{
		*product.NewProduct("P001", "Laptop Dell", 15000000, 5),
		*product.NewProduct("P002", "Chuột Logitech", 500000, 10),
		*product.NewProduct("P003", "Bàn phím cơ Keychron", 2500000, 7),
		*product.NewProduct("P004", "Màn hình LG 24 inch", 4000000, 4),
		*product.NewProduct("P005", "Tai nghe Sony", 2000000, 8),
		*product.NewProduct("P006", "Loa Bluetooth JBL", 3000000, 6),
		*product.NewProduct("P007", "Ổ cứng SSD Samsung", 1500000, 12),
		*product.NewProduct("P008", "Router WiFi TP-Link", 800000, 9),
		*product.NewProduct("P009", "MacBook Pro M1", 35000000, 3),
		*product.NewProduct("P010", "Máy in Canon", 2500000, 4),
		*product.NewProduct("P011", "Điện thoại iPhone 13", 20000000, 2),
		*product.NewProduct("P012", "Máy ảnh Sony Alpha", 45000000, 2),
		*product.NewProduct("P013", "Bàn ghế gaming", 7000000, 8),
		*product.NewProduct("P014", "Đồng hồ thông minh", 6000000, 5),
		*product.NewProduct("P015", "Máy hút bụi Xiaomi", 3000000, 10),
		*product.NewProduct("P016", "Quạt điều hòa", 4000000, 7),
		*product.NewProduct("P017", "Micro không dây", 1500000, 15),
		*product.NewProduct("P018", "Máy lọc không khí", 5500000, 4),
		*product.NewProduct("P019", "Tivi Samsung 55 inch", 12000000, 6),
		*product.NewProduct("P020", "Tai nghe không dây Bose", 8000000, 3),
	}

	var failed []string

	for _, prod := range sampleProducts {
		if err := storeInstance.Add(prod); err != nil {
			failed = append(failed, fmt.Sprintf("Mã %s: %v", prod.Code, err))
		}
	}

	fmt.Println("Dữ liệu mẫu đã được khởi tạo.")
	if len(failed) > 0 {
		fmt.Println("Một số sản phẩm không thêm được:")
		for _, errMsg := range failed {
			fmt.Println(errMsg)
		}
	} else {
		fmt.Println("Tất cả sản phẩm đã được thêm thành công.")
	}
}
