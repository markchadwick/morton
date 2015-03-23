package morton

// Encode two 8 bit unsigned integers as a single unsigned 16 bit integer. The
// first bit will be be high.
func Enc8(x, y uint8) uint16 {
	return (stripe8(x) << 1) | stripe8(y)
}

// Encode two 16 bit unsigned integers into a single 32 bit unsigned integer.
// The first parameter will be the high bit.
func Enc16(x, y uint16) uint32 {
	return (stripe16(x) << 1) | stripe16(y)
}

// Encode two 32 bit unsigned integers into a single 64 bit unsigned integers.
// The first parameter will be the high bit
func Enc32(x, y uint32) uint64 {
	return (stripe32(x) << 1) | stripe32(y)
}

// Decode two 8 bit unsigned integers from a single unsigned 16 bit interleaved.
// The high bit will go the the first return value
func Dec8(a uint16) (x, y uint8) {
	x = unstripe8(a >> 1)
	y = unstripe8(a)
	return
}

// Decode two 16 bit unsigned integers from a single unsigned 32 bit
// interleaved. The high bit will go to the first return value.
func Dec16(a uint32) (x, y uint16) {
	x = unstripe16(a >> 1)
	y = unstripe16(a)
	return
}

// Decode two 32 bit unsigned integers from a single unsigned 64 bit
// interleaved. The high bit will go to the first return value.
func Dec32(a uint64) (x, y uint32) {
	x = unstripe32(a >> 1)
	y = unstripe32(a)
	return
}

// Expands a uint8 to a uint16 so that each significant bit occupies every other
// bit of the resulting uint16.The bits are oriented such that the last bit of
// the input will be the last bit of the output. So, your high bit will be at
// bit one, not zero.
func stripe8(i uint8) (x uint16) {
	x = uint16(i)               // ---- ---- 7654 3210
	x = (x ^ (x << 4)) & 0x0f0f // ---- 7654 ---- 3210
	x = (x ^ (x << 2)) & 0x3333 // --76 --54 --32 --10
	x = (x ^ (x << 1)) & 0x5555 // -7-6 -5-4 -3-2 -1-0

	return x
}

// Stripes a 16 bit unsigned integer across 32 bits. The highest bit will be in
// position 1.
func stripe16(i uint16) (x uint32) {
	x = uint32(i)                   // ---- ---- ---- ---- fedc ba98 7654 3210
	x = (x ^ (x << 8)) & 0x00ff00ff // ---- ---- fedc ba98 ---- ---- 7654 3210
	x = (x ^ (x << 4)) & 0x0f0f0f0f // ---- fedc ---- ba98 ---- 7654 ---- 3210
	x = (x ^ (x << 2)) & 0x33333333 // --fe --dc --ba --98 --76 --54 --32 --10
	x = (x ^ (x << 1)) & 0x55555555 // -f-e -d-c -b-a -9-8 -7-6 -5-4 -3-2 -1-0

	return
}

// Stripes a 32 bit unsigned integer across 64 bits. The higest bit will be in
// position 1. Since the commends would be extremly verbose, see the above
// striping functions for explanations of the striping steps.
func stripe32(i uint32) (x uint64) {
	x = uint64(i)
	x = (x ^ (x << 16)) & 0x0000ffff0000ffff
	x = (x ^ (x << 8)) & 0x00ff00ff00ff00ff
	x = (x ^ (x << 4)) & 0x0f0f0f0f0f0f0f0f
	x = (x ^ (x << 2)) & 0x3333333333333333
	x = (x ^ (x << 1)) & 0x5555555555555555
	return
}

// Takes an interleaved 16 bit unsigned integer starting a position 1 and
// decodes it to an unsigned 8 bit integer
func unstripe8(i uint16) uint8 {
	x := i & 0x5555             // -7-6 -5-4 -3-2 -1-0
	x = (x ^ (x >> 1)) & 0x3333 // --76 --54 --32 --10
	x = (x ^ (x >> 2)) & 0x0f0f // ---- 7654 ---- 3210
	x = (x ^ (x >> 4)) & 0x00ff // ---- ---- 7654 3210

	return uint8(x)
}

// Unstripes a 32 bit unsigned integer into a 16 bit unsigned integer starting
// at the 1 position.
func unstripe16(i uint32) uint16 {
	x := i & 0x55555555             // -f-e -d-c -b-a -9-8 -7-6 -5-4 -3-2 -1-0
	x = (x ^ (x >> 1)) & 0x33333333 // --fe --dc --ba --98 --76 --54 --32 --10
	x = (x ^ (x >> 2)) & 0x0f0f0f0f // ---- fedc ---- ba98 ---- 7654 ---- 3210
	x = (x ^ (x >> 4)) & 0x00ff00ff // ---- ---- fedc ba98 ---- ---- 7654 3210
	x = (x ^ (x >> 8)) & 0x0000ffff // ---- ---- ---- ---- fedc ba98 7654 3210

	return uint16(x)
}

// Unstripes a 64 bit unsigned integer into a 32 bit unsigned integer starting
// at position 1
func unstripe32(i uint64) uint32 {
	x := i & 0x5555555555555555
	x = (x ^ (x >> 1)) & 0x3333333333333333
	x = (x ^ (x >> 2)) & 0x0f0f0f0f0f0f0f0f
	x = (x ^ (x >> 4)) & 0x00ff00ff00ff00ff
	x = (x ^ (x >> 8)) & 0x0000ffff0000ffff
	x = (x ^ (x >> 16)) & 0x00000000ffffffff

	return uint32(x)
}
