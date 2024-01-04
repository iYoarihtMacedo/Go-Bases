package main

import "fmt"

func notesAvg(values ...int) float64 {
	var sum int

	for _, value := range values {

		if value < 0 {
			return -1

		} else {
			sum += value
		}

	}

	avgResult := float64(sum) / float64(len(values))

	return float64(avgResult)

}

func main() {

	avg := notesAvg(10, 10, 8, 4, 6, 2, 0)

	if avg < 0 {

		fmt.Println("Negative values not allowed")
	} else {
		fmt.Printf("%.2f", avg)

	}

}
