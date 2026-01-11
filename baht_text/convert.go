package bahttext

import (
	"errors"
	"strings"

	"github.com/shopspring/decimal"
)

var (
	ErrNegativeNotSupported = errors.New("negative values are not supported")
)

var thaiDigits = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
var thaiUnits = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}

func ToThaiBahtText(d decimal.Decimal) (string, error) {
	if d.IsNegative() {
		return "", ErrNegativeNotSupported // เช็คติดลบ
	}

	d2 := d.Round(2) // แปลงให้เป็นทศนิยม 2 ตำแหน่ง (สตางค์) ใช้การปัดเศษ เพราะโจทย์ไม่ได้บอกให้ ตัด หรือ ปัดเศษ

	// แยก "บาท" กับ "สตางค์"
	baht := d2.Truncate(0)
	satang := d2.Sub(baht).Mul(decimal.NewFromInt(100)).Round(0) // (d2 - baht) * 100 => ส่วนทศนิยมแปลงเป็น 0–99

	bahtInt := baht.IntPart()
	satangInt := satang.IntPart()

	bahtText := numberToThaiText(bahtInt)
	if bahtText == "" {
		bahtText = thaiDigits[0]
	}

	if satangInt == 0 {
		return bahtText + "บาทถ้วน", nil
	}

	satangText := numberToThaiText(satangInt)
	if satangText == "" {
		satangText = thaiDigits[0]
	}
	return bahtText + "บาท" + satangText + "สตางค์", nil
}

func numberToThaiText(n int64) string {
	if n == 0 {
		return ""
	}

	var parts []string
	for n > 0 {
		group := n % 1000000 // เอา 6 หลักท้ายสุด, แบ่งเลขเป็นกลุ่มละ 6 หลัก (0–999,999)
		if group != 0 {
			groupText := groupToThaiText(group)
			if len(parts) > 0 { // ถ้ามีหลักล้านขึ้นไป ให้เติมคำว่า "ล้าน"
				groupText += strings.Repeat("ล้าน", len(parts))
			}
			parts = append(parts, groupText)
		} else {
			parts = append(parts, "")
		}
		n /= 1000000 // ตัด 6 หลักท้ายทิ้ง
	}

	var b strings.Builder
	for i := len(parts) - 1; i >= 0; i-- {
		if parts[i] == "" {
			continue
		}
		b.WriteString(parts[i])
	}
	return b.String()
}

func groupToThaiText(group int64) string {
	d := make([]int64, 6)
	x := group
	for i := 0; i < 6; i++ {
		d[i] = x % 10
		x /= 10
	}

	// positions: 5=แสน,4=หมื่น,3=พัน,2=ร้อย,1=สิบ,0=หน่วย
	var b strings.Builder
	for pos := 5; pos >= 0; pos-- {
		digit := d[pos]
		if digit == 0 {
			continue
		}

		switch pos {
		case 1:
			if digit == 1 {
				b.WriteString("สิบ")
			} else if digit == 2 {
				b.WriteString("ยี่สิบ")
			} else {
				b.WriteString(thaiDigits[digit])
				b.WriteString("สิบ")
			}
		case 0:
			if digit == 1 && hasHigherNonZero(d, pos) { //ถ้าหลักหน่วยเป็น 1 และ ถ้ามีหลักสิบ ให้ใช้ "เอ็ด"
				b.WriteString("เอ็ด")
			} else {
				b.WriteString(thaiDigits[digit]) // ถ้าเป็นหลักหน่วยเเดี่ยวๆ (ไม่มีเลขนำหน้า)
			}
		default:
			b.WriteString(thaiDigits[digit])
			b.WriteString(thaiUnits[pos])
		}
	}

	return b.String()
}

func hasHigherNonZero(d []int64, pos int) bool { // เช็คว่ามีหลักสิบ
	for i := pos + 1; i < len(d); i++ {
		if d[i] != 0 {
			return true
		}
	}
	return false
}
