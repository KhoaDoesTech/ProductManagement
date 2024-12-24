package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

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

		choice := ChoiceOptions()

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

func ChoiceOptions() int {
	for {
		fmt.Printf("\nChọn 1 trong %d tùy chọn: ", len(menuOptions))

		var input string
		// TODO: check error
		fmt.Scan(&input)

		trimmedInput := strings.TrimSpace(input)

		parsedChoice, err := strconv.Atoi(trimmedInput)
		if err != nil {
			fmt.Println("Lỗi::: Đầu vào không phải là số nguyên, vui lòng nhập lại.")
			continue
		}

		if parsedChoice >= 1 && parsedChoice <= len(menuOptions) {
			return parsedChoice
		}

		fmt.Println("Lỗi::: Tùy chọn không hợp lệ, vui lòng chọn lại.")
	}
}
