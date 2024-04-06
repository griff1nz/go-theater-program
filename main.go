package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var theater = [6][10]Customer{}

func theaterBuilder() {
	customers := createCustomers()
	for row := 0; row < 6; row++ {
		for col := 0; col < 10; col++ {
			customerIndex := col + (row * 10)
			if customerIndex <= len(customers)-1 {
				theater[row][col] = customers[customerIndex]
			}
		}
	}
	printTheater()
}

func createCustomers() []Customer {
	customerArr := []Customer{}
	for {
		var name string
		var height uint
		// var done string
		fmt.Println("Enter the customer's full name:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			name = scanner.Text() //Need bufio scanner to pick up whitespace https://stackoverflow.com/questions/34647039/how-to-use-fmt-scanln-read-from-a-string-separated-by-spaces
		}
		fmt.Println("Enter the customer's height in inches:")
		fmt.Scan(&height)
		for height < 1 {
			fmt.Println("Invalid height; enter a number greater than 1.")
			fmt.Scan(&height)
		}

		var newCustomer = Customer{name, height}
		fmt.Println(newCustomer)
		customerArr = append(customerArr, newCustomer)
		fmt.Printf("The customer's name is %v and they are %v inches tall.", name, height)
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n') //absorbs whitespace that is there
		if len(customerArr) >= 60 {
			fmt.Println("Theater capacity reached.")
			return customerArr
		}
		fmt.Printf("\n%v/60 customers added. Add another customer? [yes/no]", len(customerArr))
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return customerArr
		}

		// Trim any leading/trailing whitespace and convert to lowercase
		input = strings.TrimSpace(strings.ToLower(input))

		// Check user's input
		switch input {
		case "yes":
			continue
		case "no":
			return customerArr
		default:
			continue
		}
	}
}

func printTheater() {
	for _, row := range theater {
        for _, value := range row {
            // Print each value with padding
            fmt.Print(value) // Adjust padding as needed
        }
        fmt.Println() // Move to the next row
    }
}

func isSeatOccupied(row int, col int) bool {
	if theater[row][col].Height == 0 { //Height will only equal zero if a customer is not assigned
		return false
	} else {
		return true
	}
}
