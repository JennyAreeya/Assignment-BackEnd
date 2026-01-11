package bahttext

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestToThaiBahtText_Examples(t *testing.T) {
	tests := []struct {
		in   decimal.Decimal
		want string
	}{
		{decimal.NewFromFloat(1234), "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"},
		{decimal.NewFromFloat(33333.75), "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"},
	}

	for _, tc := range tests {
		got, err := ToThaiBahtText(tc.in)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if got != tc.want {
			t.Fatalf("in=%s got=%q want=%q", tc.in.String(), got, tc.want)
		}
	}
}

func TestToThaiBahtText_Zero(t *testing.T) {
	got, err := ToThaiBahtText(decimal.NewFromInt(0))
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if got != "ศูนย์บาทถ้วน" {
		t.Fatalf("got=%q want=%q", got, "ศูนย์บาทถ้วน")
	}
}

func TestToThaiBahtText_Negative(t *testing.T) {
	_, err := ToThaiBahtText(decimal.NewFromFloat(-1))
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
