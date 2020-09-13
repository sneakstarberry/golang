package main

import (
	"fmt"
	"os"
)

func main() {
	aa := os.Getenv("USERNAME")
	fmt.Println(aa)
}
