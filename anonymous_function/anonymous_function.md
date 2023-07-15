# Go lang 익명 함수
함수명을 가지지 않는 함수

익명 함수가 변수에 할당되면 변수명을 함수명처럼 사용

```
var 변수명 = func(매개변수 타입) 반환값 {
    내용
}
```
위의 코드처럼 할당하고 `변수명(매개변수)` 이런 식으로 호출할 수 있음

go에서 함수는 타입으로 취급
```
var 익명변수 = func(매개변수 타입) 반환값 {
    내용
}
함수(익명변수)

func 함수(매개변수명 func(매개변수 타입) 반환값) {
    내용
}
```
이런식으로 함수의 인자로 함수를 받는게 가능