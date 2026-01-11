# Assignments-BackEnd
Assignment Back End Stack : Golang

# Thai Baht Text Converter (Golang)
- Go 1.22+ (or recent Go version)

Convert a decimal amount (`github.com/shopspring/decimal`) into Thai Baht text with currency suffix.

## Requirements
- Input: `decimal.Decimal`
- Output: Thai text string
- Rule:
  - If the value has **no fractional part**, append **"ถ้วน"**
  - If the value has a **fractional part**, convert the fraction into Thai satang text (สตางค์)

## How to Run (Local)
go test ./...
go run .

## Project Structure

```bash
.
├── baht_text/
│   ├── convert.go
│   └── convert_test.go
├── go.mod
├── go.sum
├── main.go
└── README.md