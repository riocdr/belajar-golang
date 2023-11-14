package main

import "fmt"

func main() {

	name := [...]string{"Budi", "Ikwan", "Doni", "Anisa", "Reyhan", "Rio", "Candra"}

	slice1 := name[3:6]
	fmt.Println(slice1)
}