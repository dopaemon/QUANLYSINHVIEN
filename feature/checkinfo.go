package feature

import (
        "fmt"
        "os"
        "encoding/csv"
        _ "strconv"

	"QuanLySinhVien/models"
)

func CheckInfo(email string, isGV bool) string{
	var text string

	if isGV {
		text = "GIANGVIEN.csv"
	} else {
		text = "SINHVIEN.csv"
	}

	file, err := os.Open("./database/" + text)

	if (err != nil) {
		fmt.Println("Lỗi khi mở tệp")
		return "Error"
	}

	// Tự close file sau khi sử dụng xong :)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if (err != nil) {
		fmt.Println("Lỗi khi đọc file csv")
		return "Error"
	}

	// text := ""
	name := ""

	if isGV {
		// text := "giảng viên"
		for _, record := range records {
			userGV := models.GiangVien{
					MSGV: record[0],
					HOTEN: record[1],
					NGAYSINH: record[2],
					KHOA: record[3],
					EMAIL: record[4],
				}
			if (email == userGV.EMAIL) {
				name = userGV.HOTEN
				return name
			}
		}
	} else {
		// text := "sinh viên"
		for _, record := range records {
			userSV := models.SinhVien{
					MSSV: record[0],
					HOTEN: record[1],
					NGAYSINH: record[2],
					LOP: record[3],
					KHOA: record[4],
					EMAIL: record[5],
				}
			if (email == userSV.EMAIL) {
				name = userSV.HOTEN
				return name
			}
                }
	}

	return "Error: ?"
}
