package main

import "fmt"

// call by value
func sum(a int, b int) int {
	return a + b
}

// call by reference
func sub(a *int, b *int) {
	*a -= *b
}

// multiple return values
func div(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	a := 7
	b := 3
	fmt.Println(sum(a, b))
	sub(&a, &b)
	fmt.Println(a)
	fmt.Println(div(a, b))
}
