package main

import "fmt"

func salaryCalc(minutes int, category string) float32 {

	var salary float32
	hours := float32(minutes) / 60.0

	switch category {
	case "A":
		salary = hours * 3000
		salary *= 1.5

		break
	case "B":
		salary = hours * 1500
		salary *= 1.2

		break
	case "C":
		salary = hours * 1000

		break
	default:
		salary = -1
	}

	return salary

}

func main() {

	salary := salaryCalc(3600, "C")
	fmt.Println(fmt.Sprint("Your salary according to your category is: $", salary))
}
