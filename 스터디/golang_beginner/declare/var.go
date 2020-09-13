package main

import "fmt"

// zero value
// 여러 개 변수 선언
// 타입 생략 가능 타입
// reflect 타입 확인
// camel case 이다.
var (
	num   int
	float float32
	exam  string
)

func main() {
	num = 10
	float = 5.23
	exam = "Hello World"
	// 함수 내부에서 변수 선언
	var num2 = 20

	fmt.Println("int", num)
	fmt.Println("float", float)
	fmt.Println("string", exam)

	fmt.Println("int", num2)
}
