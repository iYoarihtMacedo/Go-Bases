package main

func main() {

	// var student1 string = 23 -> no corresponde a un String el valor asignado, sino a un int
	var student1 string

	// var grade float64 = "A" -> El valor asignado corresponde a un String en lugar de a un float64
	var grade string = "A"

	// var isEnrolled = "yes" -> Hace falta el tipo de dato, y en este caso, optaria mas por un booleano para el tipo de variable
	// Sol 1
	// var isEnrolled string = "yes"
	// Sol 2
	var isEnrolled bool = true

	// var number_of_subjects int -> No se siguen los preceptos establecidos para nombramiento de variables
	var numberOfSubjects int

	// number_of_subjects := 5 -> la variable ya se ha declarado antes, por lo que solo se tendria que asignar el valor
	numberOfSubjects = 5

}
