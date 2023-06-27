package main

import (
	"fmt"
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
	// 1
	//sort.Ints(d.elements)
	//maxDiff := d.elements[len(d.elements)-1] - d.elements[0]

	//2
	min, max := d.elements[0], d.elements[0]
	for _, v := range d.elements {
		if v < min {
			min = v
			fmt.Println("min", min)
		}
		if v > max {
			max = v
			fmt.Println("max", max)
		}
	}
	maxDiff := max - min
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
