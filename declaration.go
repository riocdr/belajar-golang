package main

import "fmt"

func main() {
	type NoKTP string
	type married bool

	var NoKtpRio NoKTP = "121251212"
	var Status married = true

	fmt.Println(NoKtpRio)
	fmt.Println(Status)
}
