# Go lang 포인터
주소값을 가지는 타입

할당된 주소에 접근해서 값을 수정할 수 있음

메모리 중복 할당을 막을 수 있음

초기화를 하지 않았을때 기본값은 nil

```
var 변수명 *타입 = 주소값

*변수명 = 변경할 값 //해당 인스턴스의 값이 변경
```

### 예제 코드
[pointer.go](./main/pointer.go)