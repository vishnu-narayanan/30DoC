package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	var j uint64
	var e float64
	var t string
	var input []string

	// Read and save an integer, double, and String to your variables.
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	j, _ = strconv.ParseUint(input[0], 10, 64)
	e, _ = strconv.ParseFloat(input[1], 64)
	t = input[2]

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + j)

	// Print the sum of the double variables on a new line.

	fmt.Printf("%0.1f\n", e+d)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

	fmt.Println(s + t)
}
