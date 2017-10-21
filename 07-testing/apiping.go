package main 

import (
	"fmt"
)

func main() {
	xs := []float64{1,2,3,4,5,6,7,8,9,10}

	avg := Average(xs)

	fmt.Println(avg)
}

func Average(xs []float64) float64 {
	total := float64(0)

	// loop into the umbers
	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}