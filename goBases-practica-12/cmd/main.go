package main

import (
	"errors"
	"fmt"
	"sort"
)

const (
	min = "minimum"
	avg = "average"
	max = "maximum"
)

func minValue(values ...float64) float64 {
	sort.Float64s(values)

	return values[0]
}

func maxValue(values ...float64) float64 {
	sort.Float64s(values)

	return values[len(values)-1]

}

func avgValue(values ...float64) float64 {
	var sum float64

	for _, value := range values {
		sum += value

	}

	return sum / float64(len(values))

}

func operation(operationType string) (func(values ...float64) float64, error) {

	if operationType == "minimum" {
		return minValue, nil
	} else if operationType == "maximum" {
		return maxValue, nil
	} else if operationType == "average" {
		return avgValue, nil
	} else {
		return nil, errors.New("no operation specified")
	}
}

func main() {
	minFunc, err := operation(min)
	minResult := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Println(minResult, err)

	maxFunc, err1 := operation(max)
	maxResult := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Println(maxResult, err1)

	avgFunc, err2 := operation(avg)
	avgResult := avgFunc(2, 3, 3, 4, 10, 2, 4, 5)
	fmt.Println(avgResult, err2)

	noFunc, err3 := operation("")
	fmt.Println(noFunc)
	fmt.Println(err3)

}
