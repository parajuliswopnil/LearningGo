package main

func fibonacciRecursive(nthItemInSeries int) int {
	var fibonacci func(n int) int
	sum := 0
	fibonacci = func(n int) int{
		if n < 2{
			return n
		}
		return fibonacci(n - 1) + fibonacci(n - 2)
	} 
	sum = fibonacci(nthItemInSeries)
	return sum
}