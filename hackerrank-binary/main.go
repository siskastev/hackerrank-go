package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// input 5 -> binary 101 -> max 1
// input 6 -> binary 110 -> max 2
// input 13 -> binary 1101 -> max 2

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	N64 := int64(n)
	binary := strconv.FormatInt(N64, 2)
	slice := strings.Split(binary, "")
	count := 0
	max := 0
	for i := range slice {
		if slice[i] == "0" {
			count = 0
		} else {
			count++
			if count > max {
				max = count
			}
		}
	}
	fmt.Println(max)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
