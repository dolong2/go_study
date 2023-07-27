package main

import (
	"fmt"
	"time"
)

func main() {
	testChan := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("보내기 전 출력")
		testChan <- 10
		time.Sleep(time.Second)
		fmt.Println("보낸후 출력")
	}()
	fmt.Println("test1")
	// 채널에서 수신할 수 있을때까지 대기
	// 즉 보낸후 go 루틴을 기다리진 않는다
	fmt.Println(<-testChan)
	fmt.Println("test2")
	time.Sleep(time.Second * 2)
}
