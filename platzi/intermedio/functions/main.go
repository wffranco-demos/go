package main

import (
	"fmt"
	"time"
)

func sum(items ...int) int {
	total := 0
	for _, item := range items {
		total += item
	}
	return total
}

func getValues(x int) (double int, triple int) {
	double = x * 2
	triple = x * 3
	return
}

func main() {
	// Testing anonymous functions
	c := make(chan int)
	go func(c chan int) {
		time.Sleep(2 * time.Second)
		c <- 1
	}(c)
	fmt.Println("Waiting for goroutine")
	<-c
	fmt.Println("Finished")

	// Testing variadic functions
	fmt.Println(sum(1, 2, 3, 4, 5))

	// Testing multiple return values
	double, triple := getValues(2)
	fmt.Println(double, triple)
}
