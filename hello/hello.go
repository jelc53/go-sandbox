package main  // go run . to run the program

import "fmt"

import "rsc.io/quote" // go mod tidy to download the package 

func main() {
	// fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}