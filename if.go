package main

import "fmt"

func main() {
	var name = "Rio"
	if name == "Rio" {
		fmt.Println("Hello Rio")
	} else {
		fmt.Println("Hi, kenalan donk")
	}
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
