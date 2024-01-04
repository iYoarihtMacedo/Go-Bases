package main

import (
	"errors"
	"fmt"
)

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	tar     = "tarantula"
)

func dogCalc(quantity float64) float64 {
	return quantity * 10

}

func catCalc(quantity float64) float64 {
	return quantity * 5

}

func hamsterCalc(quantity float64) float64 {
	return quantity * 0.250

}

func tarCalc(quantity float64) float64 {
	return quantity * 0.150

}

func Animal(animal string) (func(quantity float64) float64, error) {

	switch animal {
	case "dog":
		return dogCalc, nil

	case "cat":
		return catCalc, nil

	case "hamster":
		return hamsterCalc, nil

	case "tarantula":
		return tarCalc, nil

	default:
		return nil, errors.New("sorry, that's not an accepted animal")

	}

}

func main() {

	var amount float64

	animalDog, msg := Animal(dog)
	animalCat, msg := Animal(cat)
	animalHamster, msg := Animal(hamster)
	animalTar, msg := Animal(tar)
	notAnimal, msg := Animal("pig")

	amount += animalDog(10)
	fmt.Println(fmt.Sprint("total pet food you need: ", amount) + " Kg")
	amount += animalCat(10)
	fmt.Println(fmt.Sprint("total pet food you need: ", amount) + " Kg")
	amount += animalHamster(5)
	fmt.Println(fmt.Sprint("total pet food you need: ", amount) + " Kg")
	amount += animalTar(8)
	fmt.Println(fmt.Sprint("total pet food you need: ", amount) + " Kg")
	fmt.Println(notAnimal)
	fmt.Println(msg)

}
