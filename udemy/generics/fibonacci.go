package fibonacci

func Fibonacci[T int | int8 | int16 | int32 | int64](n T) T {
	sign := T(1)
	if n < 0 {
		n = -n
		if n%2 == 0 {
			sign = T(-1)
		}
	}
	if n == 0 || n == 1 {
		return sign * n
	}
	var a, b T = 0, 1
	var i T
	for i = 2; i <= n; i++ {
		a, b = b, a+b
		if b < 0 {
			return 0
		}
	}
	return sign * b
}
