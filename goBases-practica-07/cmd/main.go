package main

import "fmt"

// El ejercicio se podia resolver tambnien con un if sin embargo me parece que era mucho mas trabajo
// He optado por el switch ya que me ha parecido mas apropiado debido a la naturaleza del ejercicio

func main() {

	var month int

	fmt.Println("Type the number of a month")
	fmt.Scan(&month)

	switch month {
	case 1:
		fmt.Println("January")
		break
	case 2:
		fmt.Println("February")
		break
	case 3:
		fmt.Println("March")
		break
	case 4:
		fmt.Println("April")
		break
	case 5:
		fmt.Println("May")
		break
	case 6:
		fmt.Println("June")
		break
	case 7:
		fmt.Println("July")
		break
	case 8:
		fmt.Println("August")
		break
	case 9:
		fmt.Println("September")
		break
	case 10:
		fmt.Println("October")
		break
	case 11:
		fmt.Println("November")
		break
	case 12:
		fmt.Println("December")
		break

	}

}
