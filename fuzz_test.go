package morton

import (
	"github.com/google/gofuzz"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var numLoops = 1000

var _ = Describe("Fuzz Test", func() {

	fuzz := fuzz.New()

	Invariant := func(m string, f func()) {
		It(m, func() {
			for i := 0; i < numLoops; i++ {
				f()
			}
		})
	}

	Invariant("should read an 8 bit address", func() {
		var x, y uint8
		fuzz.Fuzz(&x)
		fuzz.Fuzz(&y)
		x0, y0 := Dec8(Enc8(x, y))

		Expect(x0).To(Equal(x))
		Expect(y0).To(Equal(y))
	})

	Invariant("should read an 16 bit address", func() {
		var x, y uint16
		fuzz.Fuzz(&x)
		fuzz.Fuzz(&y)
		x0, y0 := Dec16(Enc16(x, y))

		Expect(x0).To(Equal(x))
		Expect(y0).To(Equal(y))
	})

	Invariant("should read an 32 bit address", func() {
		var x, y uint32
		fuzz.Fuzz(&x)
		fuzz.Fuzz(&y)

		x0, y0 := Dec32(Enc32(x, y))

		Expect(x0).To(Equal(x))
		Expect(y0).To(Equal(y))
	})

})
