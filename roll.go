// Roll the Dice written by Ben Howell (Github/HowellBP)
// A simple Go program to test my coding skills

package main

// Imports
import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	flag.Parse()

	// Show the usage info if there are no arguments
	if len(os.Args[1:]) != 2 {
		usage()
		// Exit without error
		os.Exit(0)
	}

	// Test the arguments to see if they're numbers
	dice, ed := strconv.Atoi(flag.Arg(0))
	sides, es := strconv.Atoi(flag.Arg(1))
	// If they aren't whole numbers, throw an error and show usage info
	if ed != nil || es != nil {
		fmt.Println("Error: must use numbers")
		usage()
		// Exit with error
		os.Exit(1)
	}

	// Show the number of dice and sides per
	fmt.Printf("Rolling %vd%v\n", dice, sides)

	// Set a seed for the randomizer based on the current nanosecond
	rand.Seed(time.Now().UnixNano())

	// Set up a slice to store the result of each roll
	out := []int{}

	total := 0
	// Use the number of dice to determine the number of rolls
	for i := 1; i <= dice; i++ {
		// Store the result of each roll as a variable
		in := []int{roll(sides)}
		// Append the result to the output slice
		out = append(out, in...)
	}

	// Sum the total from all rolls and print the result
	total = sum(out...)
	fmt.Println("Total:", total)
}

// Basic usage info
func usage() {
	fmt.Println("Usage: roll [# dice] [# sides]")
	fmt.Println("Example: \"roll 1 6\"")
}

// Base the random number on the number of sides starting at 1
func roll(rolled ...int) int {
	roll := rand.Intn(rolled[0]) + 1
	return roll
}

// Add up all the results in the slice
func sum(out ...int) int {
	// Print the results of each roll
	fmt.Println("Results:", out)

	sum := 0
	// Index is discarded, output only needs the value of each roll to add to the previous value
	for _, v := range out {
		sum += v
	}
	return sum
}
