package main

type operation func(int, int) int

func add(x int, y int) int {
	return x + y
}
func sub(x int, y int) int {
	return x - y
}
func mul(x int, y int) int {
	return x * y
}
func div(x int, y int) int {
	return x / y
}

var opMap = map[string]operation{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}
