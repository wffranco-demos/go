package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int = 100

func Deposit(value int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	// balance += value
	tmp := balance
	tmp += value
	if value > 500 {
		time.Sleep(1 * time.Second)
	}
	balance = tmp
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(Balance(&lock))
}
