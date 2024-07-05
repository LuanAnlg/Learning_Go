package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getNumberOf3Digits reads a three-digit number from the user and returns its reverse
func getNumberOf3Digits() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a 3-digit number: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)

	// Ensure the input is exactly 3 digits long
	if len(input) != 3 {
		return 0, errors.New("Incompatible number, please enter exactly 3 digits")
	}

	// Parse the input as an integer
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	// Extract the digits and reverse them
	digit1 := number / 100        // Hundreds place
	digit2 := (number % 100) / 10 // Tens place
	digit3 := number % 10         // Units place

	reversedNumber := digit3*100 + digit2*10 + digit1
	return reversedNumber, nil
}

func main() {
	var reversedNumber int
	var err error

	// Loop until a valid 3-digit number is entered
	for {
		reversedNumber, err = getNumberOf3Digits()
		if err == nil {
			break
		}
		fmt.Println("Error:", err)
	}

	// Print the reversed number
	fmt.Println("Number in reverse:", reversedNumber)
}
