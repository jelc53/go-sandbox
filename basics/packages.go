package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("My favourite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)  // only capital names exported from package
	fmt.Println(add(42, 13))

}