package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// reverse array
// input N 4
// input array 1 4 3 2
// output 2 3 4 1

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	//iteration 1
	// i = 0
	// j = 4-1 = 3
	// 0 < 3 // true

	// iteration 2
	// i = 0+1 = 1
	// j = 3-1 = 2
	// 1 < 2 true

	// iteration 3
	// i = 1+1 = 2
	// j = 2-1 = 1
	// 2 < 1 false
	// stop

	// i = 0,1 -> 1,4
	// j = 3,2 -> 2,3
	//for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
	//	arr[i], arr[j] = arr[j], arr[i]
	//}

	//or
	i := 0
	j := len(arr) - 1

	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	var newArray []string
	for i := 0; i < len(arr); i++ {
		stringArray := strconv.FormatInt(int64(arr[i]), 10)
		newArray = append(newArray, stringArray)
	}
	result := strings.Join(newArray, " ")
	fmt.Println(result)

	//or
	var newArray2 []int32
	for i := len(arr) - 1; i >= 0; i-- {
		newArray2 = append(newArray2, arr[i])
	}

	fmt.Println(newArray2)
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
