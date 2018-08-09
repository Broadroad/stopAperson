package main

import (
	"log"
	"testing"
)

// TestisLessThan21 should handle negative
func TestisLessThan21(t *testing.T) {
	a := isLessThan21(-1, -2)
	if a == false {
		log.Println("error")
	}

	b := isLessThan21(0, 0)
	if b == false {
		log.Println("error")
	}

	c := isLessThan21(-1, 2)
	if c == false {
		log.Println("error")
	}

	d := isLessThan21(20, 349)
	if d == false {
		log.Println("error")
	}
}
