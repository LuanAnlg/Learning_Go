package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getNumberOf3DigitsPalindrome reads a 3-digit number from the user and checks if it is a palindrome
func getNumberOf3DigitsPalindrome() (int, error) {
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

	// Extract the hundreds and units digits
	digit1 := number / 100 // Hundreds place
	digit3 := number % 10  // Units place

	// Check if the number is a palindrome
	if digit1 != digit3 {
		return 0, errors.New("Not a palindrome number")
	}
	return number, nil
}

func main() {
	var palindrome int
	var err error

	// Loop until a valid 3-digit palindrome number is entered
	for {
		palindrome, err = getNumberOf3DigitsPalindrome()
		if err == nil {
			fmt.Println("It is a palindrome number:", palindrome)
			break
		}
		fmt.Println("Error:", err)
	}
}
