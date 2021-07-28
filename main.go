package main

import "fmt"

func SumOFElements(values []int) {

	sumOfValues := 0
	for _, score := range values {
		sumOfValues += score

	}

	fmt.Printf("The sum of the elements is = %d \n", sumOfValues)

}

func MeanValue(values []int) (int, float64) {

	sumOfValues := 0
	for _, score := range values {
		sumOfValues += score


	}

	lenOfValues := len(values)
	if lenOfValues == 0 {
		return 0, 0
	}
	meanValue := float64(sumOfValues) / float64(lenOfValues)

	return sumOfValues, meanValue

}

func main() {

	values := []int{4, 5, 1, 6, 9, 10, 32, 54, 12, 34}
	SumOFElements(values)
	fmt.Println(MeanValue(values))

}
