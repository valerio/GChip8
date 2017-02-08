package util

// CombineBytes combines two 8 bit values into a single 16 bit value.
// The high byte will be the most significant one.
func CombineBytes(low, high uint8) uint16 {
    return (uint16(high) << 8) | uint16(low)
}