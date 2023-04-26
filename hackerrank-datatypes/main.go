package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//input
//12
//4.0
//is the best place to learn and practice coding!

//output
//16
//8.0
//HackerRank is the best place to learn and practice coding!

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()

		// break the loop if line is empty
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}
	// Declare second integer, double, and String variables.
	var secondI uint64
	var secondD float64
	var secondS string

	secondI, _ = strconv.ParseUint(lines[0], 10, 64)
	secondD, _ = strconv.ParseFloat(lines[1], 64)
	secondS = lines[2]

	sumInt := i + secondI
	sumDouble := fmt.Sprintf("%.1f", d+secondD)
	concat := s + secondS

	// Print the sum of both integer variables on a new line.
	fmt.Println(sumInt)

	// Print the sum of the double variables on a new line.
	fmt.Println(sumDouble)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(concat)
}
