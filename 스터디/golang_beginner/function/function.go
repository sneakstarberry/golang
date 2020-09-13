package main

import "fmt"

// 함수 기본 식 보여주고 다중으로 변수 받고
// 반환값을 쓰는 것도 하기
// defer 써보기
func myFunc(i, j, k int) {
	fmt.Println(i, j, k)
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {

	var (
		i, j, k int
	)
	myFunc(i, j, k)
	f()
}
