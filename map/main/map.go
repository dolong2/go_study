package main

import "fmt"

func main() {
	var m1 map[int]int
	m1 = make(map[int]int)
	fmt.Println("해당 키값 조회")
	m1[10] = 20
	var value, exists = m1[10] // map에서 값을 읽어올때 첫번째 값은 value 두번째 값은 해당 키에 값이 있는지 여부를 반환
	fmt.Println(value, exists)
	fmt.Println("키값에 해당 value가 없을때")
	delete(m1, 10) // 첫번째 인자는 맵, 두번째 인자는 키값으로 해당 맵에 키값을 제거
	value, exists = m1[10]
	fmt.Println(value, exists)
	fmt.Println("for문을 사용해서 값 출력")
	m1[10] = 20
	m1[20] = 30
	m1[30] = 40
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}
