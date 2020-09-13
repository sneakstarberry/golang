package main

import (
	"fmt"
	"time"
)

// 왜 슬렉에 메시지가 안올라가지?

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func main() {
	c := make(chan int)

	go push(c)

	timer := time.After(10 * time.Second)
	tickTimerChan := time.Tick(2 * time.Second)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-timer:
			fmt.Println("timeout")
			return
		case <-tickTimerChan:
			fmt.Println("Tick")
		}
	}
}
