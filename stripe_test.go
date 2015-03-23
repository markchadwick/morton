package morton

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Integer Striping", func() {

	Context("8 bits", func() {

		It("should encode the low bit in the last position", func() {
			enc := stripe8(0xff)
			Expect(enc & 0x1).To(Equal(uint16(0x01)))
		})

		It("should stripe", func() {
			// 1010 1011 -> 0100 0100 0100 0101
			enc := stripe8(0xab)
			Expect(enc).To(Equal(uint16(0x4445)))
		})

		It("should unstripe", func() {
			dec := unstripe8(0x4445)
			Expect(dec).To(Equal(uint8(0xab)))
		})

		It("should ignore even bits", func() {
			dec := unstripe8(0x4445 | 0xaaaa)
			Expect(dec).To(Equal(uint8(0xab)))
		})

	})

	Context("16 bits", func() {

		It("should stripe", func() {
			// In:  1010 1011 1100 1101
			// Out: 0100 0100 0100 0101 0101 0000 0101 0001
			enc := stripe16(0xabcd)
			Expect(enc).To(Equal(uint32(0x44455051)))
		})

		It("should unstripe", func() {
			dec := unstripe16(0x44455051)
			Expect(dec).To(Equal(uint16(0xabcd)))
		})

	})

	Context("32 bits", func() {

		It("should stripe an int", func() {
			out := stripe32(0xffffffff)
			Expect(out).To(Equal(uint64(0x5555555555555555)))
		})

		It("should unstripe an int", func() {
			dec := unstripe32(0x5555555555555555)
			Expect(dec).To(Equal(uint32(0xffffffff)))
		})

	})

})
