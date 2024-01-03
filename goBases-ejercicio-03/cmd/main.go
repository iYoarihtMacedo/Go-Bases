package main

import (
	"fmt"
)

func main() {

	title := "Sapiens De animales a dioses. Breve historis de la humanidad"
	year := 2011
	read := false

	tense := "El libro '" + title + "' fue publicado en el año "
	tenseYear := fmt.Sprint(tense, year)

	if read {
		newTense := fmt.Sprint(tenseYear, " y ya lo lei")

		fmt.Println(newTense)
	} else {
		newTense := fmt.Sprint(tenseYear, " y no lo he leido")

		fmt.Println(newTense)
	}

}
