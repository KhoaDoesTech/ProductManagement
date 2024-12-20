# Quản lý Sản phẩm trong Cửa hàng

Xây dựng một chương trình CLI (Command Line Interface) đơn giản để quản lý sản phẩm trong một cửa hàng. Chương trình có các tính năng sau:

## 1. Tính năng chính

- **Thêm sản phẩm mới**: Người dùng nhập thông tin sản phẩm (bao gồm mã, tên, giá, số lượng) để thêm vào danh sách sản phẩm.
- **Hiển thị danh sách sản phẩm**: Hiển thị toàn bộ danh sách sản phẩm hiện có trong cửa hàng.
- **Tìm kiếm sản phẩm**: Tìm kiếm sản phẩm theo tên.
- **Xóa sản phẩm**: Xóa một sản phẩm khỏi danh sách dựa trên mã của nó.
- **Tính tổng giá trị sản phẩm**: Tính tổng giá trị (giá x số lượng) của tất cả sản phẩm trong cửa hàng.

## 2. Yêu cầu cụ thể

### a. Sản phẩm

Tạo một cấu trúc dữ liệu `Product` với các thuộc tính:

- **Name**: Tên sản phẩm (kiểu chuỗi).
- **Code**: Mã sản phẩm (kết hợp giữa kiểu chuỗi và số), ví dụ: `P0001`, `P0002`, …
- **Price**: Giá sản phẩm (kiểu `float64`).
- **Quantity**: Số lượng sản phẩm (kiểu `int`).

### b. Quản lý danh sách sản phẩm

Lưu trữ các sản phẩm trong một array hoặc slice.

### c. Hàm (Function)

- **AddProduct**: Hàm để thêm sản phẩm vào danh sách.
- **DisplayProducts**: Hàm để hiển thị danh sách sản phẩm.
- **SearchProduct**: Hàm để tìm kiếm sản phẩm theo tên.
- **DeleteProduct**: Hàm để xóa sản phẩm khỏi danh sách.
- **CalculateTotalValue**: Hàm để tính tổng giá trị sản phẩm.

### d. Giao diện (Interface)

Tạo một `StoreManager` interface để quản lý các chức năng:

- `Add(Product)`
- `Display()`
- `Search(name string) *Product`
- `Delete(name string) bool`
- `TotalValue() float64`

Cần có một struct (`Store`) thực thi interface này.

## 3. Yêu cầu thêm (Nâng cao)

- **Module**: Tách chương trình thành các module (ví dụ: `main`, `store`, `product`).
- **Xử lý lỗi**: Thêm các thông báo lỗi khi:
  - Thêm sản phẩm bị trùng tên.
  - Tìm kiếm hoặc xóa sản phẩm không tồn tại.
- **Vòng lặp**: Sử dụng vòng lặp `for` để thực hiện CLI liên tục cho đến khi người dùng chọn "Thoát".
