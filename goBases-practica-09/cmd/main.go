package main

import (
	"fmt"
)

func main() {

	var salary float64

	fmt.Println("Type your salary to calculate your tax discount: ")
	fmt.Scan(&salary)

	taxResult := tax.taxCalculator(salary)

	if taxResult != salary {
		fmt.Println(fmt.Sprint("Your salary after discount is: $", taxResult))

	} else {
		fmt.Println(fmt.Sprint("You don't have any discount in your salary: $", taxResult))

	}

}
