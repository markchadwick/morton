package morton

import (
	"strconv"
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

		It("should encode to ints", func() {
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

})

func binStr(i uint64) string {
	return strconv.FormatUint(i, 2)
}
