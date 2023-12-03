package feature

import (
	"fmt"
	"os"
	"encoding/csv"
	"strconv"
)

var (
)

           // Login, Là giảng viên
func Login() (bool, bool, string) {
	email := "default@dlu.edu.vn"
	passwd := "default"

	fmt.Print("Email: ")
	fmt.Scan(&email)

	fmt.Print("Password: " )
	fmt.Scan(&passwd)

	file, err := os.Open("./database/TAIKHOAN.csv")
	if (err != nil) {
		fmt.Println("Lỗi khi mở tệp")
		return false, false, ""
	}

	// Tự close file sau khi sử dụng xong :)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if (err != nil) {
		fmt.Println("Lỗi khi đọc file csv")
		return false, false, ""
	}

	for _, record := range records {
		emailcsv := record[0]
		passwdcsv := record[1]
		gv, _ := strconv.ParseBool(record[2])

		if (emailcsv == email && passwdcsv == passwd) {
			if (gv) {
				return true, true, email
			} else {
				return true, false, email
			}
		}
	}

	return false, false, ""
}
