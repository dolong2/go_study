# Go lang 패키지
Go는 패키지를 사용하여 코드의 모듈화, 코드의 재사용을 제공해준다.

Go는 패키지 단위로 컴포넌트를 작성하고, 이런 패키지를 활용하여 프로그램을 작성할것을 권장한다.

표준 패키지는 go를 처음 설치했던 GOROOT/pkg안에 존재

일반적인 패키지는 라이브러로서 사용되지만 main으로 명명한 경우에는 실행 프로그램으로써 인식됨

main패키지의 main메서드가 프로그램의 시작점이 된다.

다른 패키지를 가져올땐 `import`라는 키워드로 가져올 수 있다.

패키지 내에 존재하는 함수, 구조체, 인터페이스, 메서드등이 존재하는, 이들의 이름의 첫글자가 대문자여야 public으로 사용할 수 있다.

패키지내에 init이라는 메서드가 있는 경우 별도의 호출없이 실행된다.

불러올때는 `import _ '패키지명'`으로 불러올 수 있다.