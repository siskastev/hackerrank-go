package main

import (
	"fmt"
	"sort"
)

// input
// 3
// 1 2 5
// output 4
// result max 5 min 1 -> 5-1 = 4

type Difference struct {
	elements []int
}

func (d *Difference) ComputeDifference() int {
	sort.Ints(d.elements)
	maxDiff := d.elements[len(d.elements)-1] - d.elements[0]
	return maxDiff
}

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
	a := Difference{elements: elements}
	fmt.Println(a.ComputeDifference())

}
