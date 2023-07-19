package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	f1 int
	f2 string
}

func newTest(i int, str string) *Test {
	t := Test{}
	t.f1 = i
	t.f2 = str
	return &t
}

func main() {
	t1 := Test{}
	fmt.Println(t1)
	t1.f1 = 1
	t1.f2 = "Hello"
	fmt.Println(t1)
	t2 := newTest(10, "hi")
	fmt.Println(t2)
	t3 := new(Test)
	t3.f1 = 3
	t3.f2 = "Hello, world!"
	fmt.Println(t3)
	fmt.Println("t1 type == ", reflect.TypeOf(t1), "t2 type ==", reflect.TypeOf(t2), "t3 type ==", reflect.TypeOf(t3))
	// fmt.Println(t1 == t2) 타입이 달라서 에러
	// fmt.Println(t1 == t3) 타입이 달라서 에러
	fmt.Println(t2 == t3)
}
