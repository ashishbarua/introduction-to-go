package main

import "fmt"

func main() {
	customer := getData(1)
	fmt.Println(customer)
}

func getData(customerId int) (customer string) {
	firstName := "Ashish"
	lastName := "Barua"
	fullName := firstName + " " + lastName
	return fullName
}
