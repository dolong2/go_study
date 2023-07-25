package main

import "fmt"

func deferTest() {
	defer fmt.Println("defer이 적용된 출력")
	fmt.Println("defer이 적용되지 않은 출력")
}

func panicTest(a int) {
	fmt.Println("패닉되기 전")
	panic(a)
	fmt.Println("패닉된 후")
}

func recoverTest(a int) {
	defer func() {
		recover()
	}()
	panic(a)
}

func main() {
	//defer
	deferTest() // defer이 적용된 출력이 더 나중에 출력됨

	var a int
	fmt.Scan(&a)
	//recover
	recoverTest(a) // panic이 발생했지만 recover되면서 다음 문장이 실행됨

	//panic
	panicTest(a) // 패닉이 적용된 이유가 출력되고 패닉이 호출된 이후에 있는 작업은 실행되지 않음
}
