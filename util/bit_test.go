package util

import (
	"testing"
)

func TestCombineBytes(t *testing.T) {
	type args struct {
		low  uint8
		high uint8
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"combines max values", args{0xFF, 0xFF}, 0xFFFF},
		{"combines values", args{0xAB, 0xBA}, 0xBAAB},
		{"combines zero", args{0x00, 0x00}, 0x0000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombineBytes(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("CombineBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckedAdd(t *testing.T) {
	type args struct {
		a uint8
		b uint8
	}
	tests := []struct {
		name         string
		args         args
		wantResult   uint8
		wantOverflow bool
	}{
		{"add 10 and 250 overflows", args{10, 250}, 4, true},
		{"add 124 and 90 does not overflow", args{124, 90}, 214, false},
		{"add 255 and 1 overflows", args{255, 1}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotOverflow := CheckedAdd(tt.args.a, tt.args.b)
			if gotResult != tt.wantResult {
				t.Errorf("CheckedAdd() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotOverflow != tt.wantOverflow {
				t.Errorf("CheckedAdd() gotOverflow = %v, want %v", gotOverflow, tt.wantOverflow)
			}
		})
	}
}

func TestCheckedSub(t *testing.T) {
	type args struct {
		a uint8
		b uint8
	}
	tests := []struct {
		name       string
		args       args
		wantResult uint8
		wantBorrow bool
	}{
		{"subtract 10 and 250 borrows", args{10, 250}, 16, true},
		{"subtract 124 and 90 does not borrow", args{124, 90}, 34, false},
		{"subtract 0 and 1 borrows", args{0, 1}, 255, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotBorrow := CheckedSub(tt.args.a, tt.args.b)
			if gotResult != tt.wantResult {
				t.Errorf("CheckedSub() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotBorrow != tt.wantBorrow {
				t.Errorf("CheckedSub() gotBorrow = %v, want %v", gotBorrow, tt.wantBorrow)
			}
		})
	}
}
