package main

import "fmt"

func main() {
	var a int = 10

	var p1 *int = &a

	fmt.Println(a)

	*p1 = 66

	fmt.Println(a)
}
