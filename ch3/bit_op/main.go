package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // "00100010", 집합 {1, 5}
	fmt.Printf("%08b\n", y) // "00000110", 집합 {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", 교집합 {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", 집합 {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", 교집합 {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00000010", 차집합 {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // 멤버 확인
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", 집합 {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", 집합 {0, 4}
}
