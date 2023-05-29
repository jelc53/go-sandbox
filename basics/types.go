package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	var b bool
	var s string

	// default zero values for each type
	fmt.Printf("%v %v %q\n", i, b, s)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := uint(f)  // var z uint = uint(f)

	// numeric type conversions 
	fmt.Println(x, y, z)
}
