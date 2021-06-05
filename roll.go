package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var args int
var dice int
var sides int

func main() {
	args = getargs()

	flag.Parse()

	if args != 2 {
		usage()
		os.Exit(0)
	}

	dice, _ = strconv.Atoi(flag.Arg(0))
	sides, _ = strconv.Atoi(flag.Arg(1))

	fmt.Printf("Rolling %vd%v\n", dice, sides)
	rand.Seed(time.Now().UnixNano())

	total := 0
	out := []int{}
	for i := 1; i <= dice; i++ {
		in := []int{roll(sides)}
		out = append(out, in...)
	}

	total = sum(out...)

	fmt.Println("Total:", total)
}

func getargs() int {
	args = len(os.Args[1:])
	return args
}

func usage() {
	fmt.Println("Usage: roll [# dice] [# sides]")
	fmt.Println("Example: \"roll 1 6\"")
}

func roll(rolled ...int) int {
	roll := rand.Intn(rolled[0]) + 1
	return roll
}

func sum(out ...int) int {
	fmt.Println("Results:", out)

	sum := 0
	for _, v := range out {
		sum += v
	}
	return sum
}
