package main

import "fmt"

// if 문만 보여주고 else문 그리고 else if문
// 초기화 구문을 if문 안에 넣을 수 있는 것을 보여주기
// 대괄호 줄바꾸면 에러나는 거 보여주기

func main() {
	// var condition = true

	var x = "x"
	// var y = "y"
	// var z = "z"

	if condition := true; condition {
		fmt.Println(x)
	}

	// if condition == 0 {
	// 	fmt.Println(x)
	// } else if condition > 0 {
	// 	fmt.Println(y)
	// } else {
	// 	fmt.Println(z)
	// }

	// if condition
	// {
	// 	fmt.Println(x)
	// } else {
	// 	fmt.Println(y)
	// }
}
