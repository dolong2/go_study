package main

import (
	"fmt"
	"math"
)

// 인터페이스 정의
type shape interface {
	getArea() float64
	getPerimeter() float64
}

// Rect 정의
type Rect struct {
	width, height float64
}

// Circle 정의
type Circle struct {
	radius float64
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) getArea() float64 { return r.width * r.height }
func (r Rect) getPerimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	//인터페이스 사용
	r := Rect{10, 20}
	c := Circle{2}
	showArea(r, c)

	//Type Assertion
	var a interface{} = 1
	i := a
	j := a.(int)
	fmt.Println(i)
	fmt.Println(j)

	test := 12
	a = test
	i = a
	j = a.(int)
	fmt.Println(i)
	fmt.Println(j)

	f := a.(float64) // 타입에 맞지 않는 경우엔 RuntimeException
	fmt.Println(f)
}

func showArea(shapes ...shape) {
	for _, s := range shapes {
		fmt.Println(s.getArea())
	}
}
