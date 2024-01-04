package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products = []Product{
	{
		1,
		"Product1",
		10.0,
		"This is a product1",
		"Videogames",
	},
	{
		2,
		"Product2",
		20.0,
		"This is a product2",
		"Self Care",
	},
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for i := 0; i < len(Products); i++ {
		fmt.Println(Products[i])
	}

}

func getById(id int) Product {

	var p Product

	for i, prod := range Products {
		if prod.ID == id {
			p = Products[i]
		}
	}

	return p
}

func main() {

	newProd := Product{
		3,
		"Product3",
		25.0,
		"This is a new product",
		"New Category",
	}

	fmt.Println("Original Products: ")
	newProd.GetAll()

	newProd.Save()
	fmt.Println("Products Updated: ")
	newProd.GetAll()

	fmt.Println("Product Selected: ")
	productId := getById(1)
	fmt.Println(productId)
}
