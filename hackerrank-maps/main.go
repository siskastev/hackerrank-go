package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*input
3
sam 99912222
tom 11122222
harry 12299933
sam
coba
harry
*/
/*output
sam=99912222
Not found
harry=12299933
find the same
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	var mapNamePhoneNumber = make(map[string]string)
	for i := 0; i < N; i++ {
		scanner.Scan()
		split := strings.Fields(scanner.Text())
		for i := 0; i < len(split); i += 2 {
			//mapNamePhoneNumber[split[i]] = split[i+1]
			mapNamePhoneNumber[split[i]] = split[len(split)-1]
		}
	}

	//1
	var arrayName []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		arrayName = append(arrayName, line)
	}

	//or use
	//for scanner.Scan() {
	//if mapNamePhoneNumber[scanner.Text()] == "" {
	//	fmt.Printf("Not found\n")
	//} else {
	//	fmt.Printf("%v=%v\n", scanner.Text(), mapNamePhoneNumber[scanner.Text()])
	//}
	//}

	for _, v := range arrayName {
		if _, ok := mapNamePhoneNumber[v]; ok {
			fmt.Printf("%v=%v\n", v, mapNamePhoneNumber[v])
		} else {
			fmt.Printf("Not found\n")
		}
	}
}
