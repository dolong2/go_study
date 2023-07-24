package main

import "fmt"

func nextValue() func() int {
	i := 0
	return func() int {
		i++ // 바깥에 있는 변수를 익명함수내에서 사용
		return i
	}
}

func main() {
	next := nextValue()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println("----------------")

	//nextValue 함수내에 i가 선언되어있기 때문에 0부터 다시 시작
	another := nextValue()
	fmt.Println(another())
	fmt.Println(another())
	fmt.Println(another())
	fmt.Println(another())
}
