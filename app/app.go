package main

import (
	"fmt"
)

func main() {
	customers := getData()
	fmt.Println(customers)
	fmt.Println(len(customers))
}

func getData() (customers []string) {
	customers = []string{"Marcel Dempers", "Bob Smith", "John Smith"}
	customers = append(customers, "Ben Spain")
	customers = append(customers, "Aleem Janmohamed")
	customers = append(customers, "Jamie le Notre")
	customers = append(customers, "Victor Savkov")

	for _, customer := range customers {
		fmt.Println(customer)
	}
	return customers
}
