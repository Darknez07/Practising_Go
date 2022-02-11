package main

import (
	"fmt"
	"log"
)

// package for Fatal error

func main() {
	// Intialize variable
	var flt float64

	fmt.Printf("Enter the float number: ")
	// Use Scan, dynamic intialization
	_, err := fmt.Scanf("%f", &flt)
	// Error checking
	if err != nil {
		log.Fatal(err)
	}
	// Truncate to three digits
	var conv int64
	conv = int64(flt)
	fmt.Printf("The number truncated to: %d", conv)
}
