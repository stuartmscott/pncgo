package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	// Parse operation from args
	var operation func(int, int) int
	switch os.Args[1] {
	case "+":
		operation = func(a, b int) int {
			return a + b
		}
	case "-":
		operation = func(a, b int) int {
			return a - b
		}
	case "*":
		operation = func(a, b int) int {
			return a * b
		}
	case "/":
		operation = func(a, b int) int {
			return a / b
		}
	case "%":
		operation = func(a, b int) int {
			return a % b
		}
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
