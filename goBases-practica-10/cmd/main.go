package main

import "fmt"

func main() {

	// avg := notesAvg(10, 10, 8, 4, 6, 2, 0)

	// if avg < 0 {

	// 	fmt.Println("Negative values not allowed")
	// } else {
	// 	fmt.Printf("%.2f", avg)

	// }

	student := []int{10, 4, 7, 8, 10}

	for _, grade := range student {
		fmt.Println(grade)
	}

}
