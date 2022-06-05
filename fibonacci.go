package main

import (
	"fmt")
// fibonacci series with iterative method
func fibonacci(seriesLength int){
	var fiboSeries = make([]int, 2)
	fiboSeries[0] = 0
	fiboSeries[1] = 1
	for i := 0; i < seriesLength - 2; i++{
		fiboSeries = append(fiboSeries, fiboSeries[i] + fiboSeries[i+1])
	}
	fmt.Println(fiboSeries)
}

