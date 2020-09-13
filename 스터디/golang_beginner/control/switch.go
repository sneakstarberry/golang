package main

import "fmt"

// 기초로 했다가 fallthrought 이용
// 조건 없이 하는 switch 문도 이용
// 변수를 쓰지 않으면
// 첫 번째로 true인 case 조건을 수행한다.

func main() {
	i := 0

	switch i {
	case 0:
		fmt.Println("first case:", i)
		fallthrough
	case 1:
		fmt.Println("second case:", i)
	}

	switch {
	case i < 0:
		fmt.Println(i, "는 음수입니다.")
	case i == 0:
		fmt.Println(i, "는 0입니다.")
	case i > 0:
		fmt.Println(i, "는 양수입니다.")
	}
}
