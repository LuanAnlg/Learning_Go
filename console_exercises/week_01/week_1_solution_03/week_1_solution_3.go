package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculateTravelTime calculates the travel time based on distance and velocity.
func calculateTravelTime(distance, velocity float64) float64 {
	return distance / velocity
}

// getUserInput prompts the user for a numerical value using the given prompt message and returns it as a float64.
func getUserInput(prompt string) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func main() {
	// Get distance and velocity from user input
	distance, err := getUserInput("Enter distance between cities (in kilometers): ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	velocity, err := getUserInput("Enter bicycle velocity (in kilometers per hour): ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Calculate travel time using calculateTravelTime function
	travelTime := calculateTravelTime(distance, velocity)

	// Display the travel time in hours
	fmt.Printf("The travel time is %.2f hours.\n", travelTime)
}
