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
}
