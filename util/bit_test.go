package util

import "testing"

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
		{ "combines max values", args{0xFF, 0xFF}, 0xFFFF },
		{ "combines values", args{0xAB, 0xBA}, 0xBAAB },
		{ "combines zero", args{0x00, 0x00}, 0x0000 },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombineBytes(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("CombineBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
