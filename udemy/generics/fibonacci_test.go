package fibonacci

import (
	"fmt"
	"testing"
)

func runFibonacciTest[T int | int8 | int16 | int32 | int64](t *testing.T, n T, optional ...T) bool {
	var expected T = 0
	if len(optional) > 0 {
		expected = optional[0]
	}
	// t.Helper()
	result := Fibonacci(n)
	valid := (n == 0 && result == 0) || (n != 0 && ((expected == 0 && result != 0) || (expected != 0 && result == expected)))
	if valid {
		fmt.Printf("Fibonacci(%d) = %d\n", n, result)
	} else {
		t.Errorf("Fibonacci(%d) = %d; Result is invalid for type %T\n", n, result, n)
	}
	return valid
}

func TestFibonacci(t *testing.T) {
	fmt.Println("Testing Fibonacci")
	runFibonacciTest(t, 1, int16(1))
	runFibonacciTest(t, 2, int32(1))
	runFibonacciTest(t, 3, int64(2))
	runFibonacciTest(t, -1, int(1))

	// a is the maximum value for int8
	i8 := int8(1)
	for ; i8 >= 0; i8++ {
		valid := runFibonacciTest(t, i8)
		if !valid {
			break
		}
	}
	fmt.Println("Maximum fibonacci for int8 is:", Fibonacci(i8-1))

	i16 := int16(i8)
	for ; i16 > 0; i16++ {
		valid := runFibonacciTest(t, i16)
		if !valid {
			break
		}
	}
	fmt.Println("Maximum fibonacci for int16 is:", Fibonacci(i16-1))

	i32 := int32(i16)
	for ; i32 > 0; i32++ {
		valid := runFibonacciTest(t, i32)
		if !valid {
			break
		}
	}
	fmt.Println("Maximum fibonacci for int32 is:", Fibonacci(i32-1))

	i64 := int64(i32)
	for ; i64 > 0; i64++ {
		valid := runFibonacciTest(t, i64)
		if !valid {
			break
		}
	}
}
