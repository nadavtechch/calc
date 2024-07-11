package main

type calc struct {
	opMap map[string]operation
}

func (c calc) compute(x int, y int, op string) int {
	return c.opMap[op](x, y)
}
