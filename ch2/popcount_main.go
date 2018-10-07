package main

import (
	"fmt"
	"os"
	"strconv"
	"tgpl/ch2/popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		pc, err := strconv.ParseUint(arg, 10, 10)
		if err != nil {
			fmt.Fprintf(os.Stderr, "popcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(popcount.PopCount(pc))
	}
}
