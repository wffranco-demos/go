package fibonacci

import (
	"fmt"
	"math"
)

func Fibonacci(n int) int64 {
	// fibonacci for naturals between -92 and 92
	if n < -92 || n > 92 {
		fmt.Printf("Fibonacci(%d) is out of range\n", n)
		if n < -92 {
			return fibonacci7692(-92)
		}
		return fibonacci7692(92)
	}
	if n < -75 || n > 75 {
		return fibonacci7692(n)
	}
	return fibonacci75(n)
}

func fibonacci75(n int) int64 {
	// fibonacci for naturals between -75 and 75 (float64 precision)
	f := iround(powi(gr, abs(n)) / sr5)
	if n < 0 && n%2 == 0 { // if n is negative and even, we return the negative fibonacci
		return -f
	}
	return f
}

func fibonacci7692(n int) int64 {
	// fibonacci for naturals between -92 and -76, or between 76 and 92
	f := fibonacciP7692(abs(n))
	if n < 0 && n%2 == 0 { // if n is negative and even, we return the negative fibonacci
		return -f
	}
	return f
}

func fibonacciP7692(n int) int64 {
	// fibonacci for positive integers between 76 and 92
	a := fibonacci75(74)
	b := fibonacci75(75)
	for i := 76; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Fibonacci2 calculates the fibonacci number for n over 75 using recursion

func Fibonacci2(n int) int64 {
	// fibonacci for naturals between -92 and 92
	if n >= -75 && n <= 75 {
		return fibonacci75(n)
	}
	if n >= -92 && n <= 92 {
		return fibonacci76922(n)
	}
	fmt.Printf("Fibonacci(%d) is out of range\n", n)
	if n < -92 {
		return fibonacci76922(-92)
	}
	return fibonacci76922(92)
}

func fibonacci76922(n int) int64 {
	if n < 0 && n%2 == 0 { // if n is negative and even, we return the negative fibonacci
		return -fibonacciP76922(-n)
	}
	return fibonacciP76922(abs(n))
}

func fibonacciP76922(n int) int64 {
	// fibonacci for positive values between 76 and 92
	if n == 76 {
		return fibonacci75(75) + fibonacci75(74)
	}
	if n == 77 {
		return 2*fibonacci75(75) + fibonacci75(74)
	}
	return fibonacciP76922(n-1) + fibonacciP76922(n-2)
}

func Fibonacci3(n int) int64 {
	// fibonacci for naturals between -92 and 92
	if n > -93 && n < 93 {
		return fibonacci923(n)
	}
	fmt.Printf("Fibonacci(%d) is out of range\n", n)
	if n < -92 {
		return fibonacci923(-92)
	}
	return fibonacci923(92)
}

func fibonacci923(n int) int64 {
	// fibonacci for naturals between -92 and 92
	if n < 0 && n%2 == 0 { // if n is negative and even, we return the negative fibonacci
		return -fibonacciP923(-n)
	}
	return fibonacciP923(abs(n))
}

func fibonacciP923(n int) int64 {
	// fibonacci for positive values up to 92
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var a, b int64 = 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func Fibonacci4(n int) int64 {
	// fibonacci for naturals between -92 and 92
	if n >= -92 && n <= 92 {
		return fibonacci92(n)
	}
	fmt.Printf("Fibonacci(%d) is out of range\n", n)
	if n < -92 {
		return fibonacci92(-92)
	}
	return fibonacci92(92)
}

func fibonacci92(n int) int64 {
	// fibonacci for naturals between -92 and -76, or between 76 and 92
	if n < 0 && n%2 == 0 { // if n is negative and even, we return the negative fibonacci
		return -fibonacciP92(-n)
	}
	return fibonacciP92(abs(n))
}

func fibonacciP92(n int) int64 {
	// fibonacci for positive integers up to 92
	if n <= 75 {
		return fibonacciP75(n)
	}
	a := fibonacciP75(74)
	b := fibonacciP75(75)
	for i := 76; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func fibonacciP75(n int) int64 {
	return int64(math.Round(math.Pow(gr, float64(n)) / sr5))
}
