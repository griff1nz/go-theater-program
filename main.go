package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// theater := [6][10]Customer{}
	customers := createCustomers()
	fmt.Println(customers)
	fmt.Println(customers[0].Name, customers[0].Height)

}

func createCustomers() []Customer {
	customerArr := []Customer{}
	for {
		var name string
		var height uint
		var done string
		fmt.Println("Enter the customer's full name:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			name = scanner.Text() //Need bufio scanner to pick up whitespace https://stackoverflow.com/questions/34647039/how-to-use-fmt-scanln-read-from-a-string-separated-by-spaces
		}

		fmt.Println("Enter the customer's height in inches:")
		fmt.Scan(&height)

		var newCustomer = Customer{name, height}
		fmt.Println(newCustomer)
		customerArr = append(customerArr, newCustomer)
		fmt.Printf("The customer's name is %v and they are %v inches tall.\n", name, height)
		fmt.Println("Press enter to add another customer, or type 'done' if finished adding customers.")
		fmt.Scan(&done)
		if done == "done" {
			return customerArr
		}
	}
}
