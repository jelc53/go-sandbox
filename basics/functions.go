package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return  // naked return statement
}

var i, j int = 1, 2

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	// var c, python, java = true, false, "no!" // type can be omitted
	c, python, java := true, false, "no!" // short assignment statement
	fmt.Println(i, j, c, python, java)

}