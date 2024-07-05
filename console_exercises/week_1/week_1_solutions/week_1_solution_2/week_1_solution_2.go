package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculatePay calculates basic pay, gross pay, and net pay based on hours worked and hourly fee.
func calculatePay(hours, fee float64) (float64, float64, float64) {
	basic := hours * fee                    // Calculate basic pay
	gross := basic * 1.18                   // Calculate gross pay with 18% bonus
	net := gross * 0.88                     // Calculate net pay with 12% deduction
	return basic, gross, net                // Return all calculated pay values
}

// getUserInput prompts the user with a message and reads a float64 input.
func getUserInput(prompt string) (float64, error) {
	reader := bufio.NewReader(os.Stdin)     // Create a new reader to read input from standard input
	fmt.Print(prompt)                       // Print prompt message to console
	input, err := reader.ReadString('\n')   // Read input until newline
	if err != nil {
		return 0, err                       // Return error if reading input fails
	}
	input = strings.TrimSpace(input)        // Trim whitespace from input
	value, err := strconv.ParseFloat(input, 64) // Convert input to float64
	if err != nil {
		return 0, err                       // Return error if conversion fails
	}
	return value, nil                       // Return successfully parsed float64 value
}

func main() {
	// Obtain hours worked and hourly fee from user input
	hours, err := getUserInput("Enter hours worked: ")
	if err != nil {
		fmt.Println("Error: ", err)         // Print error message if input error occurs
		return
	}

	fee, err := getUserInput("Enter hourly fee: ")
	if err != nil {
		fmt.Println("Error: ", err)         // Print error message if input error occurs
		return
	}

	// Calculate pay using calculatePay function
	basic, gross, net := calculatePay(hours, fee)

	// Display the calculated pay values
	fmt.Printf("Basic pay: %.2f\n", basic)
	fmt.Printf("Gross pay: %.2f\n", gross)
	fmt.Printf("Net pay: %.2f\n", net)
}
