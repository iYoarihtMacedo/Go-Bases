package main

import "fmt"

func taxCalculator(salary float64) float64 {

	if salary > 50000.0 {
		return salary * (0.83)

	} else if salary > 150000.0 {
		return salary * (0.73)

	} else {
		return salary
	}

}

func main() {

	var salary float64

	fmt.Println("Type your salary to calculate your tax discount: ")
	fmt.Scan(&salary)

	taxResult := taxCalculator(salary)

	if taxResult != salary {
		fmt.Println(fmt.Sprint("Your salary after discount is: $", taxResult))

	} else {
		fmt.Println(fmt.Sprint("You don't have any discount in your salary: $", taxResult))

	}

}
