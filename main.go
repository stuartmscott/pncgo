package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Parse operation from args
	operation, err := ParseOperation(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// Parse parameters from args
	a, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	// Calculate result
	result := operation(a, b)
	// Print result
	log.Println(result)
}

func ParseOperation(op string) (func(int, int) int, error) {
	switch op {
	case "+":
		return func(a, b int) int {
			return a + b
		}, nil
	case "-":
		return func(a, b int) int {
			return a - b
		}, nil
	case "*":
		return func(a, b int) int {
			return a * b
		}, nil
	case "/":
		return func(a, b int) int {
			return a / b
		}, nil
	case "%":
		return func(a, b int) int {
			return a % b
		}, nil
	default:
		return nil, fmt.Errorf("Unrecognized operation: %s", op)
	}
}
