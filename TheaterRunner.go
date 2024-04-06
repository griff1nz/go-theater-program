package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	theaterBuilder()
	reader := bufio.NewReader(os.Stdin)
	outerLoop:
	for {
		fmt.Println("What would you like to do?\n1. Randomize the seating arrangement\n2. Check if a seat is occupied\n3. Find the most occupied row\n4. Find the tallest customer\n5. Find customers to be moved\n6. Reserve two seats\n7. Exit")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
		}

		// Trim any leading/trailing whitespace and convert to lowercase
		input = strings.TrimSpace(input)

		// Check user's input
		switch input {
		case "1":
			randomizeSeats()
		case "2":
			{
				var row int
				var col int
				fmt.Println("Enter the seat's row (1-6):")
				fmt.Scan(&row)
				for row > 6 || row < 1 {
					fmt.Println("Invalid input; enter a valid row number. ")
					fmt.Scan(&row)
				}
				fmt.Println("Enter the seat's column (1-10):")
				fmt.Scan(&col)
				for col > 10 || row < 1 {
					fmt.Println("Invalid input; enter a valid column number. ")
					fmt.Scan(&col)
				}
				if isSeatOccupied(row, col) {
					fmt.Println("This seat is occupied.")
					reader.ReadString('\n')
				} else if !isSeatOccupied(row, col) {
					fmt.Println("This seat is available. ")
					reader.ReadString('\n') //Absorbing whitespace
				}
			}
		case "3":
			fmt.Println("The most occupied row is",findMostOccupiedRow())
		case "4":
			fmt.Println("The tallest customer in the theater is", getTallestCustomer())
		case "5":
			fmt.Println("The following customers are seated behind someone more than 3 inches taller than them and need to be moved:")
			for i := 0; i < len(getCustomersToBeMoved()); i++ {
				fmt.Println(getCustomersToBeMoved()[i].Name)
			}
		case "6":
			reserveTwoSeats()
		case "7":
			fmt.Println("Bye bye!")
			break outerLoop
		default:
			fmt.Println("Invalid input.")
			continue
		}
	}

}
