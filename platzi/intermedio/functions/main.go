package main

import (
	"fmt"
	"time"
)

// Testing anonymous functions
func main() {
	c := make(chan int)
	go func(c chan int) {
		time.Sleep(2 * time.Second)
		c <- 1
	}(c)
	fmt.Println("Waiting for goroutine")
	<-c
	fmt.Println("Finished")
}
