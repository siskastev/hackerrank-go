package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Bad String")
	} else {
		fmt.Println(input)
	}

}
