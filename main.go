package main

import "fmt"

func main() {
	c := calc{}
	x := 10
	y := 5
	op := "+"
	result := c.compute(x, y, op)
	fmt.Println(result)
}
