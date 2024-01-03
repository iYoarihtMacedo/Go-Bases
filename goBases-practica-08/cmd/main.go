package main

import "fmt"

func main() {

	var employees = map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	var oldestThan21 int

	for employee, age := range employees {
		employeeData := fmt.Sprint("Employee Name: "+employee+", Employee Age: ", age)
		fmt.Println(employeeData)

		if age > 21 {
			oldestThan21 += 1
		}

	}

	fmt.Println(fmt.Sprint("Employees oldest than 21: ", oldestThan21))
	employees["Federico"] = 25
	delete(employees, "Pedro")

	for employee, age := range employees {
		employeeData := fmt.Sprint("Employee Name: "+employee+", Employee Age: ", age)
		fmt.Println(employeeData)

	}
}
