package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Expense struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
}

var expenses []Expense

func main() {
	// Menu chính
	for {
		fmt.Println("\n*** Quản lý Chi tiêu Cá nhân ***")
		fmt.Println("1. Thêm Khoản Chi tiêu")
		fmt.Println("2. Xem Danh Sách Chi tiêu")
		fmt.Println("3. Tính Tổng Chi tiêu")
		fmt.Println("4. Lưu vào File")
		fmt.Println("5. Tải từ File")
		fmt.Println("6. Thoát")
		fmt.Print("Chọn một tùy chọn (1-6): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addExpense()
		case 2:
			listExpenses()
		case 3:
			total := totalExpenses()
			fmt.Printf("Tổng Chi tiêu: %.2f\n", total)
		case 4:
			saveToFile("expenses.json")
		case 5:
			loadFromFile("expenses.json")
		case 6:
			fmt.Println("Thoát ứng dụng...")
			os.Exit(0)
		default:
			fmt.Println("Tùy chọn không hợp lệ. Vui lòng chọn lại.")
		}
	}
}

func addExpense() {
	var amount float64
	var description, date string

	fmt.Print("Nhập số tiền: ")
	fmt.Scan(&amount)
	fmt.Print("Nhập mô tả: ")
	fmt.Scan(&description)
	fmt.Print("Nhập ngày (yyyy-mm-dd): ")
	fmt.Scan(&date)

	expenses = append(expenses, Expense{Amount: amount, Description: description, Date: date})
	fmt.Println("Khoản chi tiêu đã được thêm thành công!")
}

func listExpenses() {
	if len(expenses) == 0 {
		fmt.Println("Chưa có khoản chi tiêu nào.")
		return
	}

	fmt.Println("\nDanh Sách Chi tiêu:")
	for i, expense := range expenses {
		fmt.Printf("%d. Số tiền: %.2f, Mô tả: %s, Ngày: %s\n", i+1, expense.Amount, expense.Description, expense.Date)
	}
}

func totalExpenses() float64 {
	total := 0.0
	for _, expense := range expenses {
		total += expense.Amount
	}
	return total
}

func saveToFile(filename string) {
	data, err := json.Marshal(expenses)
	if err != nil {
		fmt.Println("Lỗi khi lưu vào file:", err)
		return
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Lỗi khi ghi vào file:", err)
	} else {
		fmt.Println("Dữ liệu đã được lưu vào file", filename)
	}
}

func loadFromFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Lỗi khi đọc từ file:", err)
		return
	}
	err = json.Unmarshal(data, &expenses)
	if err != nil {
		fmt.Println("Lỗi khi giải mã dữ liệu từ file:", err)
	} else {
		fmt.Println("Dữ liệu đã được tải từ file", filename)
	}
}
