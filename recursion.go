package main

func factorialRecursive(factorialOf int)int{
	factorial := 0
	if factorialOf == 0{
		return 1
	}
	factorial = factorialOf * factorialRecursive(factorialOf - 1)
	return factorial
}