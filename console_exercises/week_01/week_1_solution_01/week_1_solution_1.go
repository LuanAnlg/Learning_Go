package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const pi float64 = 3.14159

func semicircleArea(radius float64) float64 {
	return (pi * radius * radius) / 2
}

func triangleArea(radius, side float64) float64 {
	return 0.5 * radius * math.Sqrt(side*side - radius*radius)
}

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
	radius, err := getUserInput("Enter the radius of the circle: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	side, err := getUserInput("Enter the side of the triangle: ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	figureArea := semicircleArea(radius) + triangleArea(radius, side)
	fmt.Println("Figure area:", figureArea)
}
