# Morton
Efficient encoding and decoding of [Z-Order
Curves](http://en.wikipedia.org/wiki/Z-order_curve) to and from a euclidean
grid. God, that sounds pompous as shit. You probably know if you needed this
thing. It's in Go.


## Example
```go
package main

import (
	"log"
	"math/rand"

	"github.com/markchadwick/morton"
)

func main() {
	// First, pick two random numbers and log them to stdout
	x0 := uint32(rand.Int31())
	y0 := uint32(rand.Int31())
	log.Printf("x0:   %d", x0)
	log.Printf("y0:   %d", y0)

	// Encode these numbers as a single index on a Z-ordered curve
	addr := morton.Enc32(x0, y0)
	log.Printf("addr: %d", addr)

	// From that index on the Z-ordered curve, find an x, y coordinate and print
	// it to stdout
	x1, y1 := morton.Dec32(addr)
	log.Printf("x1:   %d", x1)
	log.Printf("y1:   %d", y1)
}
```

```
x0:   1298498081
y0:   2019727887
addr: 3882731965896198231
x1:   1298498081
y1:   2019727887
```
