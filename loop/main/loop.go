package main

import "fmt"

func main() {
	var input int
	fmt.Scan(&input)

	//1. 기본
	for i := 0; i < input; i++ {
		fmt.Println(i)
	}
	fmt.Println("---------------------------")

	//2. 생략한 for
	var count = 10
	for ; count >= 0; count-- {
		fmt.Println(count)
	}
	fmt.Println("---------------------------")

	//3. 조건식만 쓴 for
	for input >= 0 {
		input--
	}
	fmt.Println(input)
	fmt.Println("---------------------------")

	//4. 무한 for문
	for {
		fmt.Println("infinite loop!!!")
	}

}
