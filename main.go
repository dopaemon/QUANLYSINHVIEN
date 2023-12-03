package main

import (
	_ "encoding/csv"
	"fmt"
	"os"

	_ "QuanLySinhVien/models"
	"QuanLySinhVien/feature"
)

var (
	LoginActive	bool
	isGV	bool
)

func CheckLogin(check bool) bool {
	fmt.Print("Bạn có tài khoản chưa ?\n[y/N]: ")

	input := "Biến Input"

	for ; ; {
		fmt.Scan(&input)

		if (input == "y" || input == "Y") {
			LoginActive, isGV = feature.Login()
			if (!LoginActive) {
				fmt.Println("Đăng nhập thất bại, bạn muốn thử lại ?")
				fmt.Print("[y/N]: ")
				fmt.Scan(&input)
				if (input == "y" || input == "Y") {
					return CheckLogin(true)
				} else {
					return false
				}
			} else {
				fmt.Println("Đăng nhập thành công !!!")
				return true
			}
		} else if (input == "n" || input == "N") {
			fmt.Println("Liên lạc giảng viên để thêm tài khoản !!!")
			return false
		} else {
			fmt.Println("Chọn [y/N]: ")
		}
	}

	return false
}

func main() {
	if (!CheckLogin(true)) { os.Exit(0) }
}
