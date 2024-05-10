# Go lang 인터페이스
Go lang에서 인터페이스는 메서드의 집합이다.

하나의 인터페이스를 구현할려면 단순히 해당 메서드를 구현하면 됨

인터페이스는 구조체와 마찬가지로 type문을 사용하여 선언

```
type 인터페이스이름 interface {
    함수명() 반환값
}
```

Go에서 종종 빈 인터페이스를 자주 접하게 되는데, 이를 인터페이스 타입이라고도 부른다.

`interface{}`와 같이 표현할 수 있다.

인터페이스 타입은 메서드를 0개를 구현하므로 어떤 타입이든 담을 수 있음

java의 Object와 같다라고 생각하면 됨

### Type Assertion
interface type의 x와 타입 T에 대하여 x.(T)로 표현할때, 이는 x가 nil이 아니며,

x는 T타입에 속한다는 점을 확인하는 것으로 이러한 표현을 "Type Assertion"이라 부름

### 예제 코드
[interface.go](./main/interface.go)