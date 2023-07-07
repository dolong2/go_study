package main

import "fmt"

func main() {
	// 1
	var num1 int
	num1 = 1
	fmt.Println(num1)

	// 2
	var num2, num3, num4 int
	num2, num3, num4 = 2, 3, 4
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)

	// 3
	var str1 = "test"
	fmt.Println(str1)

	// 4
	var (
		str2 string
		num5 int
	)
	str2 = "test2"
	num5 = 5
	fmt.Println(str2)
	fmt.Println(num5)

	// 5
	test := "test3"
	fmt.Println(test)
}
