# Go lang 변수
변수를 선언하고 사용하지 않으면 에러남

일반적으론 `var 변수명 타입`형식으로 정의

같은 타입의 변수를 여러개 선언할땐`var 변수명1, 변수명2, 변수명3 타입` 이런 방식으로 선언

변수를 선언과 동시에 초기화할땐 타입을 지정하지 않아도 됨 ex).`var a = 1`

다른 타입의 변수를 여러개 선언할땐 `var (변수명1 타입 변수명2 타입)`

함수안에선 `변수명 := 값`이런 방식으로 var를 생략할 수 있음 -> 함수 밖에선 동작하지 않음

## Go lang 변수 타입
* uint8 -> 음수제외 8비트 정수
* uint16 -> 음수제외 16비트 정수
* uint32 -> 음수제외 32비트 정수
* uint64 -> 음수제외 64비트 정수
* uint -> 32비트에선 uint32, 64비트에선 uint64
* int8 -> 8비트 정수
* int16 -> 16비트 정수
* int32 -> 32비트 정수
* int64 -> 64비트 정수
* int -> 32비트에선 int32, 64비트에선 int64
* float32 -> 32비트 부동 소수점, 7자리
* float64 -> 64비트 부동 소수점, 12자리
* complex64 -> float32 크기의 복소수
* complex128 -> float64 크기의 복소수
* uintptr -> uint와 같은 크기의 포인터
* bool -> 참/거짓을 표현하는 8비트 자료형
* byte -> 8비트 자료형
* rune -> 유니코드를 저장하는 자료형, int32와 크기는 동일
* string -> 문자열을 저장하기위한 자료형

### 예제 코드
[variable.go](./main/variable.go)