package main

import (
	"fmt"
	"strings"
)

//input "Go time"
//output :
//Goo
//tiimmmeeee

func main() {
	msg := "Go time"
	repeat(msg)
}

func repeat(msg string) {
	words := strings.Split(msg, " ")
	for _, v := range words {
		var repeat []string
		for i, s := range v {
			word := strings.Repeat(string(s), i+1)
			repeat = append(repeat, word)
		}
		fmt.Println(strings.Join(repeat, ""))
	}
}
