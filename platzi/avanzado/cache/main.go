package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	value any
	err   error
}
type IntResult struct {
	value int
	err   error
}

/*********************
 * No cached Section *
 *********************/
func Fibonacci(n int) int {
	if n < 0 || n > 92 {
		return -1
	}
	if n <= 1 {
		return n
	}
	a := 0
	b := 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Recursive Fibonacci
func RFibonacci(n int) int {
	if n < 0 || n > 42 {
		return -1
	}
	if n <= 1 {
		return n
	}
	return RFibonacci(n-1) + RFibonacci(n-2)
}

/*************************************************************
 * Cache 1 Section - Caching the current value and the error *
 *************************************************************/
type Memory1 struct {
	cache map[int]Result
}

func (m *Memory1) GetFibonacci(n int) (any, error) {
	// using the recursive function
	r := RFibonacci(n)
	if r < 0 {
		return nil, fmt.Errorf("number out of range")
	}
	return r, nil
}

func NewCache1() *Memory1 {
	return &Memory1{
		cache: make(map[int]Result),
	}
}

func (m *Memory1) Fibonacci(n int) (any, error) {
	res, ok := m.cache[n]
	if !ok {
		res.value, res.err = m.GetFibonacci(n)
		m.cache[n] = res
	}
	return res.value, res.err
}

/*************************************************************
 * Cache 2 Section - Caching the current value and the error *
 *************************************************************/
type Memory2 struct {
	cache map[int]Result
}

func (m *Memory2) GetFibonacci(n int) (any, error) {
	// using the non recursive function
	return Fibonacci(n), nil
}

func NewCache2() *Memory2 {
	return &Memory2{
		cache: make(map[int]Result),
	}
}

func (m *Memory2) Fibonacci(n int) (any, error) {
	res, ok := m.cache[n]
	if !ok {
		res.value, res.err = m.GetFibonacci(n)
		m.cache[n] = res
	}
	return res.value, res.err
}

/*********************************************************
 * Cache 3 Section - Recursive caching values and errors *
 *********************************************************/
type Memory3 struct {
	cache map[int]IntResult
}

func NewCache3() *Memory3 {
	return &Memory3{
		cache: map[int]IntResult{0: {0, nil}, 1: {1, nil}},
	}
}

func (m *Memory3) Fibonacci(n int) (int, error) {
	if n < 0 || n > 92 {
		return 0, fmt.Errorf("invalid number")
	}
	res, ok := m.cache[n]
	if !ok {
		f1, err := m.Fibonacci(n - 2)
		if err != nil {
			return 0, err
		}
		f2, err := m.Fibonacci(n - 1)
		if err != nil {
			return 0, err
		}
		r := f1 + f2
		if r < 0 {
			m.cache[n] = IntResult{0, fmt.Errorf("number is too big")}
		} else {
			m.cache[n] = IntResult{r, nil}
		}
		res = m.cache[n]
	}
	return res.value, res.err
}

/***************************************************************************
 * Cache 3 Section - Recursive caching values. No recursive function calls *
 ***************************************************************************/
type Memory4 struct {
	cache map[int]int
	last  int
	lock  sync.Mutex
}

func NewCache4() *Memory4 {
	return &Memory4{
		cache: map[int]int{0: 0, 1: 1},
		last:  1,
	}
}

func (m *Memory4) Fibonacci(n int) int {
	if n < 0 || n > 92 {
		return -1
	}
	m.lock.Lock()
	if n <= m.last {
		m.lock.Unlock()
		return m.cache[n]
	}
	a := m.cache[m.last-1]
	b := m.cache[m.last]
	for i := m.last + 1; i <= n; i++ {
		a, b = b, a+b
		m.cache[i] = b
	}
	m.last = n
	m.lock.Unlock()
	return m.cache[n]
}

/*******************
 * Main Section
 *******************/
func main() {
	f1 := NewCache1()
	f2 := NewCache2()
	f3 := NewCache3()
	f4 := NewCache4()
	fibo := []int{5, 10, 15, 20, 20, 25, 30, 30, 35, 40, 40, 45, 50, 60, 70, 80, 90, 92}
	for _, v := range fibo {
		var duration time.Duration
		start0 := time.Now()
		var value any = Fibonacci(v)
		duration = time.Since(start0)
		fmt.Printf("Fibonacci(%d) = %d, Time: %v\n", v, value, duration)
		start1 := time.Now()
		value, _ = f1.Fibonacci(v)
		duration = time.Since(start1)
		fmt.Printf("Fibonacci1(%d) = %d, Time: %v, Memory: %d\n", v, value, duration, len(f1.cache))
		start2 := time.Now()
		value, _ = f2.Fibonacci(v)
		duration = time.Since(start2)
		fmt.Printf("Fibonacci2(%d) = %d, Time: %v, Memory: %d\n", v, value, duration, len(f2.cache))
		start3 := time.Now()
		value, _ = f3.Fibonacci(v)
		duration = time.Since(start3)
		fmt.Printf("Fibonacci3(%d) = %d, Time: %v, Memory: %d\n", v, value, duration, len(f3.cache))
		start4 := time.Now()
		value = f4.Fibonacci(v)
		duration = time.Since(start4)
		fmt.Printf("Fibonacci4(%d) = %d, Time: %v, Memory: %d\n", v, value, duration, len(f4.cache))
		fmt.Println()
	}
}
