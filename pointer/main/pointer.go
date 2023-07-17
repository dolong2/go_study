package main

import "fmt"

func main() {
	var a int = 10

	var p1 *int

	fmt.Println(p1)

	p1 = &a

	fmt.Println(p1)

	fmt.Println(a)

	*p1 = 66

	fmt.Println(a)
}
