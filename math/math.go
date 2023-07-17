//reference: https://www.youtube.com/watch?v=XQzTUa9LPU8

package math

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Divide(x, y int) float64 {
	if y == 0 {
		return 0.0
	}
	return float64(x / y)
}

func Multiply(x, y int) int {
	return x * y
}

//ref : https://medium.com/code-chasm/unit-test-in-go-8d17d3910818

func Factorial(n uint) uint {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}
