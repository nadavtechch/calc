package main

type calc struct {
}

func (c calc) compute(x int, y int, op string) int {
	return opMap[op](x, y)
}
