package main

import (
	"fmt"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

var c, python, java bool

func main() {

	var i int
	fmt.Println(i, c, python, java)

	fmt.Println("My favorite number is", rand.Intn(10))
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("The sum of %d and %d is %d \n", x, y, add(x, y))

	a, b := swap("First", "Second")

	fmt.Println(a, b)
	fmt.Println(Split(17))
	PrintRandom()
	PrintExported()
}
