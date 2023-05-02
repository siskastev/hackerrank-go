package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

//input -> output
//array 4
//3 5 -> 243, 2 4 ->16
//-1 2 -> n and p should be non-negative, 2 -2 -> n and p should be non-negative

type Calculator struct {
	n, p int
}

func NewCalculator(n int, p int) *Calculator {
	return &Calculator{n: n, p: p}
}

func (c *Calculator) power() {
	if c.n < 0 || c.p < 0 {
		fmt.Println("n and p should be non-negative")
	} else {
		pow := math.Pow(float64(c.n), float64(c.p))
		fmt.Println(int(pow))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < N; i++ {
		var (
			n, p int
		)
		fmt.Scanln(&n, &p)
		c := NewCalculator(n, p)
		c.power()
	}
}
