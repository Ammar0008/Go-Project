package main

import (
	"fmt"
)

func showMenu() {
	fmt.Println("Welcome to the calculator")
	fmt.Println("Choose an operation:")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("5. Exit the calculator")
}

func getUserInput() (int, float64, float64) {
	var choice int
	var num1, num2 float64

	fmt.Print("Enter the number of operation (1-5): ")
	fmt.Scan(&choice) // Removed quotes around &choice (was a string earlier)

	if choice >= 1 && choice <= 4 { // Corrected the conditional and added a missing closing brace
		fmt.Print("Enter the first number: ")
		fmt.Scan(&num1)
		fmt.Print("Enter the second number: ")
		fmt.Scan(&num2)
	}
	return choice, num1, num2 // Corrected missing commas and fixed return statement
}

func calculate(choice int, num1, num2 float64) float64 { // Changed Calculate to lowercase to match Go convention
	switch choice {
	case 1:
		return num1 + num2
	case 2:
		return num1 - num2
	case 3:
		return num1 * num2
	case 4:
		if num2 != 0 { // Moved this inside the case
			return num1 / num2
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
			return 0
		}
	default:
		return 0
	}
}

func main() {
	for {
		showMenu()
		choice, num1, num2 := getUserInput()

		if choice == 5 {
			fmt.Println("Exiting the calculator.")
			break
		} else if choice >= 1 && choice <= 4 {
			result := calculate(choice, num1, num2) // Function name corrected
			fmt.Printf("Result: %.2f\n\n", result)
		} else {
			fmt.Println("Invalid choice. Please try again.\n")
		}
	}
}
