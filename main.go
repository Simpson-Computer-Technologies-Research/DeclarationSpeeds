package main

import (
	"fmt"
	"time"
)

// Main Function
func main() {
	PreDeclaredType()
	NoDeclaredType()
}

// Pre declare the variables type
func PreDeclaredType() {
	// Declare a start time for tracking speed
	var startTime time.Time = time.Now()

	// Declare the s variable 10^8 amount of times
	for i := 0; i < 1000000000; i++ {
		var s string = "this is a string!"
		_ = s
	}

	// Print the speed
	fmt.Printf("\nPreDeclaredType() -> %v\n", time.Since(startTime))
}

// No declared variable type
func NoDeclaredType() {
	// Declare a start time for tracking speed
	var startTime time.Time = time.Now()

	// Declare the s variable 10^8 amount of times
	for i := 0; i < 1000000000; i++ {
		s := "this is a string!"
		_ = s
	}

	// Print the speed
	fmt.Printf("\nNoDeclaredType() -> %v\n\n", time.Since(startTime))
}
