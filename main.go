package main

import (
	"fmt"

	"github.com/shopspring/decimal"

	bahttext "Assignment-Backend/baht_text"
)

// ตัว runner ไว้เแสดงผลลัพธ์ตัวอย่างการแปลงจำนวนเงินเป็นข้อความภาษาไทย
func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(-0),
		decimal.NewFromFloat(-33333.75),
	}

	for _, input := range inputs {
		fmt.Println("Input:", input.String())
		out, err := bahttext.ToThaiBahtText(input) // เรียกใช้ function logic แปลงจำนวนเงินเป็นข้อความภาษาไทย
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Println("Output:", out)
		fmt.Println("----")
	}
}
