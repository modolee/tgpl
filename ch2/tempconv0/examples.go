package tempconv0

import "fmt"

func Example1() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100"
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180"
	//fmt.Printf("%g\n", boilingF-tempconv0.FreezingC) // 컴파일 오류 : 타입 불일치
}

func Example2() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	//fmt.Println(c == f) // 컴파일 오류 : 타입 불일치
	fmt.Println(c == Celsius(f)) // "true"!
}

func Example3() {
	c := FToC(212.0)
	fmt.Println(c.String()) // "100ºC"
	fmt.Printf("%v\n", c)   // "100ºC"
	fmt.Printf("%s\n", c)   // "100ºC"
	fmt.Println(c)          // "100ºC"
	fmt.Printf("%g\n", c)   // "100"; String을 호출하지 않음
	fmt.Println(float64(c)) // "100"; String을 호출하지 않음
}
