package utils

import (
	"fmt"
	"os"
)

var (
)

func ContentMenu_Home(isGV bool, tex string, name string) {
	text := ""

	fmt.Print("Đăng nhập với tư cách ", tex)
	Echo(" " + name, true)

	fmt.Println("\nChọn 1 chức năng: ")

	if isGV {
		text1 := " [1]: Thông tin của tôi\n"
		text2 := " [2]: Thêm sinh viên\n"
		text3 := " [3]: Xoá sinh viên\n"
		text4 := " [4]: Sửa sinh viên\n"
		text5 := " [5]: Tìm sinh viên\n"
		text6 := " [0]: Đăng xuất\n"
		text = text1 + text2 + text3 + text4 + text5 + text6
	}
	fmt.Print(text)

	fmt.Print(" Nhập [0/5]: ")
}

func Menu_Choose(input int) {
	temp := ""
	switch(input) {
		case 1:
			fmt.Println("Chức năng 1")
			fmt.Scan(&temp)
			break
		case 2:
			fmt.Println("Chức năng 2")
			fmt.Scan(&temp)
			break
		case 3:
			fmt.Println("Chức năng 3")
			fmt.Scan(&temp)
			break
		case 4:
			fmt.Println("Chức năng 4")
			fmt.Scan(&temp)
			break
		case 5:
			fmt.Println("Chức năng 6")
			fmt.Scan(&temp)
			break
		default:
			fmt.Println("Chưa chọn chức năng !!!")
			fmt.Scan(&temp)
			break
	}
}

func Menu(isGV bool, text string, name string) int {
	input := -1
	for ; ; {
		ClearScreen()
		ContentMenu_Home(isGV, text, name)
		fmt.Scan(&input)
		if (input == 0) {
			Echo("\nĐã đăng xuất, Tạm biệt " + text + " " + name + " !!!", true)
			os.Exit(0)
		}
		Menu_Choose(input)
	}
	return input
}
