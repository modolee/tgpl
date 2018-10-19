package main

import "fmt"

func main() {
	var s1 = []string{"h", "e", "l", "l", "o"}
	s2 := "hello, world"

	s1[0] = "H"
	fmt.Println(s1)
	fmt.Println(s2)
}
