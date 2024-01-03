package main

import "fmt"

func main() {

	var age int
	var seniority float64
	var salary float64

	fmt.Println("Type your age")
	fmt.Scan(&age)
	fmt.Println("type your seniority at work")
	fmt.Scan(&seniority)
	fmt.Println("type your salary")
	fmt.Scan(&salary)

	fmt.Println("\n...\n")

	if age < 22 {
		fmt.Println("Sorry! You are not old enough to get a credit")
	} else if seniority < 1.0 {
		fmt.Println("Sorry! Your seniority at work is not enough to get a credit")
	} else if salary < 100000.0 {
		fmt.Println("Great! You could get a credit paying an interest")
	} else {
		fmt.Println("Congratulations! You could get a free interest credit")
	}

}
