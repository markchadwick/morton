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
