package morton

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Morton Suite")
}

var _ = Describe("Encoding", func() {

	Context("8 bits", func() {

		It("should encode two ints", func() {
			var a uint8 = 0xff // 1111 1111
			var b uint8 = 0x00 // 0000 0000

			enc := Enc8(a, b) // 1010 1010 1010 1010
			Expect(enc).To(Equal(uint16(0xAAAA)))
		})

		It("should decode two ints", func() {
			x, y := Dec8(0x5e0b)             // 0101 1110 0000 1011
			Expect(x).To(Equal(uint8(0x33))) // 0011 0011
			Expect(y).To(Equal(uint8(0xe1))) // 1110 0001
		})

	})

	Context("16 bits", func() {

		It("should encode two ints", func() {
			var a uint16 = 0xabcd // 1010 1011 1100 1101
			var b uint16 = 0xef01 // 1110 1111 0000 0001

			// Expect 1101 1100 1101 1111 1010 0000 1010 0011
			enc := Enc16(a, b)
			Expect(enc).To(Equal(uint32(0xDCDFA0A3)))
		})

		It("should decode two ints", func() {
			x, y := Dec16(0xDCDFA0A3)
			Expect(x).To(Equal(uint16(0xabcd)))
			Expect(y).To(Equal(uint16(0xef01)))
		})

	})

	Context("32 bits", func() {

		It("should encode two ints", func() {
			var a uint32 = 0x01234567
			var b uint32 = 0x89abcdef

			enc := Enc32(a, b)
			Expect(enc).To(Equal(uint64(0x40434C4F70737C7F)))
		})

		It("should decode two ints", func() {
			x, y := Dec32(0x40434C4F70737C7F)
			Expect(x).To(Equal(uint32(0x01234567)))
			Expect(y).To(Equal(uint32(0x89abcdef)))
		})

	})

})

func binStr(i uint64, bits int) string {
	s := strconv.FormatUint(i, 2)
	return strings.Repeat("0", bits-len(s)) + s
}
