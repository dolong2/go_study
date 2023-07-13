package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [3]int
	arr2 := [...]int{3, 4, 5, 6}
	var arr3 [4]int
	fmt.Println("각 배열의 값")
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println("각 배열의 타입")
	arr1Type := reflect.TypeOf(arr1)
	arr2Type := reflect.TypeOf(arr2)
	arr3Type := reflect.TypeOf(arr3)
	fmt.Println("arr1 =", arr1Type)
	fmt.Println("arr2 =", arr2Type)
	fmt.Println("arr3 =", arr3Type)
	fmt.Println("각 배열의 타입 비교")
	fmt.Println("arr1Type == arr2Type ->", arr1Type == arr2Type)
	fmt.Println("arr2Type == arr3Type ->", arr2Type == arr3Type)
	fmt.Println("arr1Type == arr3Type ->", arr1Type == arr3Type)
	fmt.Println("각 배열의 값 비교")
	//fmt.Println("arr1 == arr2 ->", arr1 == arr2) -> 타입이 일치하지 않아서 에러
	fmt.Println("arr2 == arr3 ->", arr2 == arr3) // 배열 원소가 일치하지 않아서 false
	//fmt.Println("arr1 == arr3 ->", arr1 == arr3) -> 타입이 일치하지 않아서 에러

	var slice1 []int
	var slice2 = make([]int, 4)
	var slice3 = []int{1, 2, 3, 4}
	fmt.Println("각 슬라이스의 값")
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println("각 슬라이스의 타입")
	slice1Type := reflect.TypeOf(slice1)
	slice2Type := reflect.TypeOf(slice2)
	slice3Type := reflect.TypeOf(slice3)
	fmt.Println("arr1 =", slice1Type)
	fmt.Println("arr2 =", slice2Type)
	fmt.Println("arr3 =", slice3Type)
	fmt.Println("각 슬라이스의 타입 비교")
	fmt.Println("slice1Type == slice2Type ->", slice1Type == slice2Type)
	fmt.Println("slice2Type == slice3Type ->", slice2Type == slice3Type)
	fmt.Println("slice1Type == slice3Type ->", slice1Type == slice3Type)
	fmt.Println("슬라이스는 == 연산자로 비교할 수 없다.")
}
