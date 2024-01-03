package main

import "fmt"

func main() {

	var word string

	fmt.Println("Write a word:")
	fmt.Scan(&word)

	fmt.Println(word)
	fmt.Println("Word size: ", len(word))
	fmt.Println("Letters: ")

	for i := 0; i < len(word); i++ {
		fmt.Println(string(word[i]))

	}

}
