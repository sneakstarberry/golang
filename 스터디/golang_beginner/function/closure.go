package main

import (
	"fmt"
)

// func makeSuffix(suffix string) func(string) string {
// 	return func(name string) string {
// 		if !strings.HasSuffix(name, suffix) {
// 			return name + suffix
// 		}
// 		return name
// 	}
// }

// func main() {
// 	addZip := makeSuffix(".zip")
// 	addTgz := makeSuffix(".tar.gz")
// 	fmt.Println(addTgz("go1.5.1.src"))
// 	fmt.Println(addZip("go1.5.1.windows-amd64"))
// }
func callback(y int, f func(int, int)) {
	f(y, 2) // add(1, 2)를 호출
}

func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b) // 1 + 2 = 3
}

func main() {
	callback(1, add) // 1 + 2 = 3
}
