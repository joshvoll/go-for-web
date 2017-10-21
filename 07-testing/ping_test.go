package main

import (
	"testing"
)

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1,2,3,4,5,6,7,8,9,10})

	if v != 5.5 {
		t.Error("Expected 5.5 got: ", v)
	}
}