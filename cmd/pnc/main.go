package main

import (
	"github.com/stuartmscott/pncgo/expression"
	"log"
	"os"
)

func main() {
	// Parse expression from args
	exp, err := expression.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// Calculate result
	result, err := exp.Resolve()
	if err != nil {
		log.Fatal(err)
	}
	// Print result
	log.Println(result)
}
