package morton

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Morton Suite")
}

var _ = Describe("Encoding bytes", func() {

	Context("when expanding 8 to 16 bits", func() {

		It("should encode the low bit in the last position", func() {
			enc := enc8To16(0xff)
			Expect(enc & 0x1).To(Equal(uint16(0x01)))
		})

		It("should interleave bits", func() {
			enc := enc8To16(0xff)
			Expect(enc).To(Equal(uint16(0x5555)))
		})

		It("should interleave bytes", func() {
			a := 0xff // 1111 1111
			b := 0x00 // 0000 0000
		})

		It("should interleave two bytes", func() {
			a := 0x33 // 00110011
			b := 0xe1 // 11100001
		})
	})

})

// var _ = Describe("Encoding 2 bytes to 16 bits", func() {
//
// 	It("should encode a high X", func() {
// 		enc := Enc2To16(0xff, 0x00)
// 		log.Printf("Encoded: %08x", enc)
// 		Expect(enc).To(Equal(0xaaaa))
// 	})
//
// 	// It("should encode a high Y", func() {
// 	// 	enc := Enc2To16(0x00, 0xff)
// 	// 	Expect(enc).To(Equal(0x5555))
// 	// })
//
// })
