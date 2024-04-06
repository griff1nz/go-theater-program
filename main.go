package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var theater = [6][10]Customer{}
var customers = []Customer{}

func theaterBuilder() { //Assigns customers in row-major order by default
	customers = createCustomers()
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

func createCustomers() []Customer { //Creates customers from user input
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

func randomizeSeats() { //Randomly assigns the customers to seats in the theater
	for row := 0; row < 6; row++ {
		for col := 0; col < 10; col++ {
			theater[row][col].Height = 0
			theater[row][col].Name = "" //Initialize theater
		}
	}
	for i := 0; i < len(customers); {
		row := rand.Intn(6)
		col := rand.Intn(10)
		if theater[row][col].Height == 0 {
			theater[row][col] = customers[i]
			i++
		} else {
			continue
		}
	}
	printTheater()
}

func printTheater() { //Prints the theater as an organized table
	for _, row := range theater {
		for _, value := range row {
			// Print each value with padding
			fmt.Print(value) // Adjust padding as needed
		}
		fmt.Println() // Move to the next row
	}
}

func isSeatOccupied(row int, col int) { //Tells if the seat is occupied or not
	if theater[row][col].Height == 0 { //Height will only equal zero if a customer is not assigned
		fmt.Println("This seat is available")
	} else {
		fmt.Println("This seat is occupied")
	}
}

func findMostOccupiedRow() { //Gets the most occupied row
	var currentRowOccupancy int
	var maxRowOccupancy int
	for row := 0; row < 6; row++ {
		for col := 0; col < 10; col++ {
			if theater[row][col].Height != 0 {
				currentRowOccupancy++
			}
		}
		if currentRowOccupancy > maxRowOccupancy {
			maxRowOccupancy = currentRowOccupancy
		}
	}
	fmt.Println("The most occupied row is", maxRowOccupancy)
}

func getTallestCustomer() { //Gets the tallest customer in the theater
	var tallest Customer
	for row := 0; row < 6; row++ {
		for col := 0; col < 10; col++ {
			if theater[row][col].Height > tallest.Height {
				tallest = theater[row][col]
			}
		}
	}
	fmt.Println("The tallest customer in the theater is", tallest.Name)
}

func getCustomersToBeMoved() []Customer { //Returns customers that need to be moved due to someone more than 3 inches taller than them being seated in front of them
	var customersToBeMoved []Customer
	for row := 1; row < 6; row++ { //Only need to do rows 1-5 because nobody is in front of row 0
		for col := 0; col < 10; col++ {
			if theater[row-1][col].Height-3 > theater[row][col].Height {
				customersToBeMoved = append(customersToBeMoved, theater[row][col])
			}
		}
	}
	return customersToBeMoved
}

func reserveTwoSeats() {
	fmt.Println("Not implemented yet")
}
