package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

var p Point

func main() {
	var heads, tails int

	switch coinflip() {
	case "heads":
		heads++
		fmt.Printf("heads : %d\n", heads)
	case "tails":
		tails++
		fmt.Printf("tails : %d\n", tails)
	default:
		fmt.Println("landed on edge!")
	}

	fmt.Println(Signum(3))
	fmt.Println(Signum(-3))
}

func coinflip() string {
	return "tails"
}

func Signum(x int) int {
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
