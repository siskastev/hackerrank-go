package main

import (
	"fmt"
	"strings"
)

//input
// 2
//Hacker
//Rank

//output
//Hce akr
//Rn ak

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var N int
	var S string

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&S)
		even, odd := changeString(S)
		fmt.Printf("%v %v\n", even, odd)
	}
}

func changeString(s string) (string, string) {
	split := strings.Split(s, "")
	var even []string
	var odd []string
	for i := 0; i < len(split); i++ {
		if i%2 == 0 {
			even = append(even, split[i])
		}
	}

	for i := 0; i < len(split); i++ {
		if i%2 == 1 {
			odd = append(odd, split[i])
		}
	}

	stringEven := strings.Join(even, "")
	stringOdd := strings.Join(odd, "")

	return stringEven, stringOdd
}
