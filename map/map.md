# Go lang map
key - value 형식으로 저장할 수 있는 자료구조

`var 변수명 map[Key타입]value타입` 이렇게 선언할 수 있다.

이때 선언된 map은 nil값을 가지고, 이를 초기화하기 위해선 make함수를 사용한다.

`맵이름 = make(map[Key타입]value타입)` <- 이렇게 초기화할 수 있다.

make말고 리터럴을 사용해서 초기화 할 수 있다.

`맵이름 := map[key타입]value타입{key값: value값}`

### 예제 코드
[map.go](./main/map.go)