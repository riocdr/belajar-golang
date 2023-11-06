package main

import "fmt"

func main() {
	/*	const firstname string = "Rio"
		const lastname = "Candra"
		const age = 29

		fmt.Println(firstname)
		fmt.Println(lastname)
		fmt.Println(age)
	*/
	//dapat disingkat seperti dibawah//

	const (
		firstname string = "Rio"
		lastname         = "Candra"
		age              = 29
	)
	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(age)

}
