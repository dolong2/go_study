package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)

	// if문
	if s < 10 {
		fmt.Println("10미만인 수")
	} else if s >= 10 && s < 20 {
		fmt.Println("10 이상 20 미만")
	} else {
		fmt.Println("그외 다른수")
	}

	//switch/case문
	switch s {
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	default:
		fmt.Println("10과 20이 아닌 다른수")
	}
}
