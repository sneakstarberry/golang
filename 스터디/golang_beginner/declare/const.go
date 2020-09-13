package main

import "fmt"

// 전역 상수 설정
// 타입을 정하지 않아도 자동으로 상수의 타입을 결정
const (
	sunday   = iota // 0
	monday          // 1
	tuesday         // 2
	thursday        // 3
	friday          // 4
	saturday        // 5
)

func main() {
	fmt.Println(sunday)
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(thursday)
	fmt.Println(friday)
	fmt.Println(saturday)
}
