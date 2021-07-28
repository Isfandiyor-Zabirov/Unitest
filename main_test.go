package main

import "testing"

func TestMeanValue(t *testing.T) {

	tests := []struct {
		name              string
		values            []int
		SumOfElementsWant int
		MeanValueWant     float64
	}{
		{"When all elements are positive", []int{4, 5, 1, 6, 9, 10, 32, 54, 12, 34}, 167, 16.7},
		{"When all elements are negative", []int{-4, -5, -1, -6, -9, -10, -32, -54, -12, -34}, -167, -16.7},
		{"When elements are positive and negative", []int{4, -5, -1, 6, 9, -10, -32, 54, 12, 34}, 71, 7.1},
	}

	for _, test := range tests {
		SumOFElements, meanValue := MeanValue(test.values)

		if SumOFElements != test.SumOfElementsWant || meanValue != test.MeanValueWant {
			t.Errorf("MeanValue test %s got SumofElements %d SumofElementswant %d got meanValue %.2f meanValuewant %.2f", test.name, SumOFElements, test.SumOfElementsWant, meanValue, test.MeanValueWant)
		}
	}

}
