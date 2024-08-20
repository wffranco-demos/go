package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int = 100

func Deposit(value int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	// balance += value
	tmp := balance
	tmp += value
	time.Sleep(1 * time.Second)
	balance = tmp
	lock.Unlock()
}

func Balance() int {
	return balance
}

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance())
}
