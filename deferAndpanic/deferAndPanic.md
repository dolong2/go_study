# Go lang defer panic
## defer
defer 키워드는 특정 문장 혹은 함수를 해당 함수가 return 하기 전에 실행하는 키워드

타 언어의 finally 블록처럼 clean-up 작업을 위해 사용
## panic
Go 내장 함수로써 현재 함수를 즉시 종료하고 현재 함수의 defer문을 전부 실행하고, 리턴한다

계속 콜스택을 타고 올라가다 프로그램이 에러를 내고 종료하게 된다.
## recover
해당 함수는 panic함수에 의한 상태를 다시 정상상태로 되돌리는 함수

### 예제 코드
[defer.go](./main/defer.go)