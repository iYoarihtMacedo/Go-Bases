package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (emp Employee) PrintEmployee() {
	fmt.Println("Employee data: ")
	fmt.Printf("Employee ID: %d\n", emp.ID)
	fmt.Println("Employee Position: ", emp.Position)
	fmt.Printf("Employee Person ID: %d\n", emp.Person.ID)
	fmt.Println("Employee Name: ", emp.Name)
	fmt.Println("Employee Date of Birth: ", emp.DateOfBirth)

}

func main() {

	yoariht := Person{
		1,
		"Yoariht Macedo",
		"30/01/1998",
	}

	brenda := Person{
		2,
		"Brenda Martinez",
		"19/07/1998",
	}

	emp1 := Employee{
		1,
		"Cybersecurity Consultant",
		brenda,
	}

	emp2 := Employee{
		2,
		"Software Developer",
		yoariht,
	}

	emp1.PrintEmployee()
	emp2.PrintEmployee()

}
