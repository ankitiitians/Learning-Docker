package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calculator <number1> <operator> <number2>")
		os.Exit(1)
	}

	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error parsing number1:", err)
		os.Exit(1)
	}

	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Error parsing number2:", err)
		os.Exit(1)
	}

	operator := os.Args[2]

	var result float64
	var resultErr error

	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "/":
		result, resultErr = divide(num1, num2)
	default:
		fmt.Println("Invalid operator:", operator)
		os.Exit(1)
	}

	if resultErr != nil {
		fmt.Println("Error:", resultErr)
		os.Exit(1)
	}

	fmt.Printf("%v %s %v = %v\n", num1, operator, num2, result)
}
