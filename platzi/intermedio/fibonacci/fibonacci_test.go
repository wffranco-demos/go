package fibonacci

import (
	"testing"
)

func runFibonacciTest(t *testing.T, n int, expected int64) bool {
	// t.Helper()
	got1 := Fibonacci(n)
	if got1 != expected {
		t.Errorf("Fibonacci(%d) = %d; expected %d", n, got1, expected)
	}
	got2 := Fibonacci(n)
	if got2 != expected {
		t.Errorf("Fibonacci2(%d) = %d; expected %d", n, got2, expected)
	}
	got3 := Fibonacci(n)
	if got3 != expected {
		t.Errorf("Fibonacci3(%d) = %d; expected %d", n, got3, expected)
	}
	return got1 == expected && got2 == expected && got3 == expected
}

func TestFibonacci(t *testing.T) {
	runFibonacciTest(t, 0, 0)
	runFibonacciTest(t, 1, 1)
	runFibonacciTest(t, 2, 1)
	runFibonacciTest(t, 3, 2)
	runFibonacciTest(t, 4, 3)
	runFibonacciTest(t, 5, 5)
	runFibonacciTest(t, 6, 8)
	runFibonacciTest(t, 7, 13)
	runFibonacciTest(t, 8, 21)
	runFibonacciTest(t, 9, 34)
	runFibonacciTest(t, 10, 55)
	runFibonacciTest(t, 20, 6765)
	runFibonacciTest(t, 30, 832040)
	runFibonacciTest(t, 40, 102334155)
	runFibonacciTest(t, 50, 12586269025)
	runFibonacciTest(t, 60, 1548008755920)
	runFibonacciTest(t, 70, 190392490709135)
	runFibonacciTest(t, 80, 23416728348467685)
	runFibonacciTest(t, 90, 2880067194370816120)
}

func TestFibonacciFull(t *testing.T) {
	if !runFibonacciTest(t, 0, 0) {
		return
	}
	if !runFibonacciTest(t, 1, 1) {
		return
	}

	var a int64 = 0
	var p int64 = 1 // positive fibonacci
	var n int64     // negative fibonacci
	for i := 2; i <= 92; i++ {
		a, p = p, a+p
		if i%2 == 0 {
			n = -p
		} else {
			n = p
		}
		if !runFibonacciTest(t, i, p) || !runFibonacciTest(t, -i, n) {
			return
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Fibonacci(30)
		Fibonacci(75)
		// Fibonacci(92)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Fibonacci2(30)
		Fibonacci2(75)
		// Fibonacci2(92)
	}
}

func BenchmarkFibonacci3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Fibonacci3(30)
		Fibonacci3(75)
		// Fibonacci3(92)
	}
}

func BenchmarkFibonacci4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Fibonacci4(30)
		Fibonacci4(75)
		// Fibonacci4(92)
	}
}
