package main

import "fmt"

type Student struct {
	Name     string
	LastName string
	DNI      int
	Date     string
}

func (s *Student) details(name string, lastName string, dni int, date string) {
	(*s).Name = name
	(*s).LastName = lastName
	(*s).DNI = dni
	(*s).Date = date

}

func main() {

	var studentName string
	var studentLastName string
	var studentDni int
	var studentsDate string

	fmt.Println("Welcome Student, please type your data")
	fmt.Println("Name: ")
	fmt.Scan(&studentName)
	fmt.Println("Last Name: ")
	fmt.Scan(&studentLastName)
	fmt.Println("DNI: ")
	fmt.Scan(&studentDni)
	fmt.Println("Enters Date: ")
	fmt.Scan(&studentsDate)

	newStudent := Student{
		"",
		"",
		0,
		"",
	}

	newStudent.details(
		studentName,
		studentLastName,
		studentDni,
		studentsDate,
	)

	fmt.Println("\n\nNice, your student details are: ")
	fmt.Println(newStudent)

}
