package main

import "fmt"

func main() {

	var name [3]string

	name[0] = "Reyhan"
	name[1] = "Andaru"
	name[2] = "Candra"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var umur = [3]int{
		25,
		10,
		8,
	}
	fmt.Println(umur)
	fmt.Println(umur[0])
	fmt.Println(umur[1])
	fmt.Println(umur[2])
}
