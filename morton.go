package morton

func Enc2To16(x, y uint8) uint16 {
	return 0
}

// Expands a uint8 to a uint16 so that each significant bit occupies every other
// bit of the resulting uint16.The bits are oriented such that the last bit of
// the input will be the last bit of the output. So, your high bit will be at
// bit one, not zero.
func enc8To16(i uint8) uint16 {
	x := uint16(i)

	x = (x ^ (x << 8)) & 0x00ff
	x = (x ^ (x << 4)) & 0x0f0f
	x = (x ^ (x << 2)) & 0x3333
	x = (x ^ (x << 1)) & 0x5555

	return x
}

func Dec2From16(a uint16) (uint8, uint8) {
	return 0, 0
}
