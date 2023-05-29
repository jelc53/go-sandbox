package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	sum := 0
	// stanfard for loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// init and post are optional
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// while loop (drop ;)
	for sum < 10000 {
		sum += sum
	}
	fmt.Println(sum)

	// if statement function
	fmt.Println(sqrt(2), sqrt(-4))
}