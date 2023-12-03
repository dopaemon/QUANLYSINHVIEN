package main

import (
	_ "encoding/csv"
	"fmt"
	_ "os"

	"QuanLySinhVien/models"
)

func main() {
	user1 := models.GiangVien{MSGV: 1, HOTEN: "Nguyễn Văn A", NGAYSINH: "1/1/1970", KHOA: "CNTT", EMAIL: "1@dlu.edu.vn" }
	fmt.Println("MSGV: ", user1.MSGV)
	fmt.Println("HOTEN: ", user1.HOTEN)
	fmt.Println("NGAYSINH: ", user1.NGAYSINH)
	fmt.Println("KHOA: ", user1.KHOA)
	fmt.Println("EMAIL: ", user1.EMAIL)
}
