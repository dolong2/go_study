package main

import (
	"fmt"
	"time"
)

func routineTest() {
	time.Sleep(time.Second * 2)
	fmt.Println("2초 대기후 출력")
}

func main() {
	routineTest()
	fmt.Println("test1") // go루틴을 적용하지 않아서 나중에 출력

	go routineTest()
	fmt.Println("test2") // go루틴을 적용해서 먼저 출력

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("go루틴 익명함수 출력")
	}()

	func() {
		fmt.Println("일반 익명함수 출력")
	}()

	time.Sleep(time.Second * 3)
}
