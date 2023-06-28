package main

import "fmt"

//input n = 5 give array 1 4 4 3 9
//output can see the sun is 1 4 4 9 then count there are 4

func main() {
	var (
		N        int
		elements []int
	)
	fmt.Scanln(&N)
	for i := 0; i < N; i++ {
		var tempElement int
		fmt.Scanf("%d", &tempElement)
		elements = append(elements, tempElement)
	}
	fmt.Println("Can see the sun:", calculate(elements))
}

func calculate(arr []int) int {
	count := 1
	maxHeight := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] >= maxHeight {
			count++
			maxHeight = arr[i]
		}
	}

	return count
}
