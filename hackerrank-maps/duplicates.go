package main

import (
	"fmt"
	"sort"
)

// input N = 5 => 4 3 3 2 2
// output [2,3]
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

	sort.Ints(elements)
	temp := make(map[int]bool)
	result := []int{}
	for i := 0; i < len(elements); i++ {
		if _, ok := temp[elements[i]]; ok {
			result = append(result, elements[i])
		}
		temp[elements[i]] = true
	}

	fmt.Println(result)

}
