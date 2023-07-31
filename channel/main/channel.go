package main

import (
	"fmt"
	"time"
)

func sendChan(ch chan<- string) {
	ch <- "Data"
	// x := <-ch // 에러발생
}

func receiveChan(ch <-chan string) {
	// ch <- "Data" 에러발생
	data := <-ch
	fmt.Println(data)
}

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

	// 다른 go루틴에서 주고 받지 않아서 데드락
	// channel := make(chan int)
	// channel <- 1 // 수신하는 루틴이 없어서 데드락
	// fmt.println(<- channel)// 수신해도 데드락 (별도의 Go루틴없기 때문)

	// 버퍼 채널로 만들어서 최대 버퍼수까지 데이터를 보낼 수 있어서 에러X
	ch := make(chan int, 1)
	ch <- 101
	fmt.Println(<-ch)

	ch2 := make(chan string, 1)
	sendChan(ch2)
	receiveChan(ch2)

	ch3 := make(chan int, 2)
	ch3 <- 1
	ch3 <- 2
	close(ch3) // 채널 닫기
	// ch3 <- 3 채널이 닫혔는데 송신하면 오류
	for i := range ch3 {
		fmt.Println(i)
	}
}
