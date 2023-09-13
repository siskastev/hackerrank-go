package main

import (
	"fmt"
	"regexp"
)

func main() {
	var str string

	fmt.Scanln(&str)

	regex := regexp.MustCompile("[A-Z][a-z]*")
	words := regex.FindAllString(str, -1)
	count := 1 + len(words)

	//or
	//count := 1
	//for _, r := range str {
	//	if unicode.IsUpper(r) {
	//		count++
	//	}
	//}

	fmt.Println("Number of words found:", count)

}
