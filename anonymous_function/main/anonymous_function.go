package main

import "fmt"

func main() {
	// 익명함수
	sum := func(a int, b int) int {
		return a + b
	}
	fmt.Println(sum(1, 2))

	//함수를 파라미터로
	sub := func(a int, b int) int {
		return a - b
	}
	calc := func(f func(int, int) int, a int, b int) int {
		return f(a, b)
	}
	fmt.Println(calc(sum, 1, 2))
	fmt.Println(calc(sub, 1, 2))
}
