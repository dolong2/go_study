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

type test struct {
	a int
	b int
}

func (t test) sum() int {
	return t.a + t.b
}

func main() {
	a := 7
	b := 3
	fmt.Println(sum(a, b))
	sub(&a, &b)
	fmt.Println(a)
	fmt.Println(div(a, b))
	var t = test{a: 10, b: 20}
	fmt.Println(t.sum())
}
