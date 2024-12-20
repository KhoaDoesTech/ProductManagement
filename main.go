package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/KhoaDoesTech/ProductManagement/actions"
)

type MenuOption struct {
	Message string
	Action  actions.Action
}

var menuOptions []MenuOption

func init() {
	menuOptions = []MenuOption{
		{Message: "Thêm sản phẩm mới", Action: actions.AddProduct{}},
		{Message: "Hiển thị danh sách sản phẩm", Action: actions.DisplayProducts{}},
		{Message: "Tìm kiếm sản phẩm theo tên", Action: actions.SearchProduct{}},
		{Message: "Xóa sản phẩm theo mã", Action: actions.DeleteProduct{}},
		{Message: "Tính tổng giá trị sản phẩm", Action: actions.CalculateTotalValue{}},
		{Message: "Khởi tạo dữ liệu mẫu", Action: actions.InitData{}},
		{Message: "Thoát", Action: actions.ExitApplication{}},
	}
}

func main() {
	for {
		ClearScreen()
		fmt.Println("\n=== Quản lý Sản phẩm trong Cửa hàng ===")
		for index, option := range menuOptions {
			fmt.Printf("%d. %s\n", index+1, option.Message)
		}

		var choice int
		for {
			fmt.Printf("\nChọn 1 trong %d tùy chọn: ", len(menuOptions))
			_, err := fmt.Scan(&choice)

			if err == nil && choice >= 1 && choice <= len(menuOptions) {
				break
			}

			fmt.Println("Lỗi::: Tùy chọn không hợp lệ, vui lòng chọn lại.")
		}

		ClearScreen()
		menuOptions[choice-1].Action.Execute()
		PauseScreen()
	}
}

func PauseScreen() {
	fmt.Println("\nNhấn Enter để tiếp tục...")
	fmt.Scanln()
	fmt.Scanln()
}

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear") // Linux/Unix/MacOS
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}