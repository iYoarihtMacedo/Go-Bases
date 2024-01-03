package main

import "fmt"

func main() {

	lat := 19.638674
	long := -99.104407
	cd := "Coacalco de Berriozabal"

	tense := "La ciudad '" + cd + "' tiene una latitud de "
	tense1 := " y una longitud de "

	tenseLat := fmt.Sprint(tense, lat)
	tenseLong := fmt.Sprint(tense1, long)

	fmt.Println(tenseLat + tenseLong)

}
