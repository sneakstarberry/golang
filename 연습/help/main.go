package main

import "fmt"

func main() {
	a := "party"

	for _, s := range a {
		fmt.Sprintf("%s\n", s)
	}
}
