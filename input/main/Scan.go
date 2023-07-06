package main

import "fmt"

func main() {
	var userName string
	fmt.Scan(&userName) // 입력받을때 주소값, 입력받을 갯수만큼 매개변수에 넣어주기, 반환값은 입력성공한 매개변수 갯수, 줄바꿈(엔터)도 입력받지 않아서 새줄에서 입력해도 됨
	fmt.Println(userName + "님 안녕하세요")
	fmt.Scanf("%s", &userName) // 입력받을때 주소값, c scanf처럼 사용, 반환값은 입력성공한 매개변수 갯수
	fmt.Println(userName + "님 안녕하세요")
	var testStr1, testStr2 string
	fmt.Scanln(&testStr1, &testStr2) // 한줄에서 공백을 기준으로 나눠서 입력, 반환값은 입력성공한 매개변수 갯수
	fmt.Println(testStr1, testStr2)
}
