package main

import "fmt"

//input 6
// output 1+2+3+6
//input n 25 -> output 1+5+25 = 31

type AdvancedArithmatic interface {
	DivisorSum(n int) int
}

type Calculator struct {
	n int
}

func (c *Calculator) DivisorSum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			total += i
		}
	}
	return total
}

func main() {
	var N int
	fmt.Scanln(&N)
	c := Calculator{N}
	//var arithmetic AdvancedArithmatic
	//arithmetic = &c
	arithmetic := AdvancedArithmatic(&c)
	fmt.Println(arithmetic.DivisorSum(c.n))
}
