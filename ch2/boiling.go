// Boiling은 물의 끓는점을 출력합니다.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gºF or %gºC\n", f, c)
	// 출력:
	// boiling pint = 212 ºF or 100ºC
}
