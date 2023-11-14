package main

import "fmt"

func main() {

	person := map[string]string{

		"name" : "Rio",
		"address" : "Unit2",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := map[string]string{}
	book["title"] = "Belajar Golang"
	book["author"] = "Rio Candra"
	book["wrong"] = "Ups"

	fmt.Println(book)
	delete(book,"wrong")

	fmt.Println(book)
}