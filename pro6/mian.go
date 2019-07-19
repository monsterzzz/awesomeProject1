package main

import (
	"fmt"
)

func main() {
	fmt.Println("pro6")
	// 运算符
	a := 20
	b := 30
	c := 5
	fmt.Println(c)
	c++
	fmt.Println(c)
	c--
	fmt.Println(c)
	fmt.Println(a+b, a-b, b-a, a*b, a/b, b/a)

	// fmt.Println(a && b) error
	var d = true
	var e = false

	fmt.Println(d && d, d && e)
	fmt.Println(d || d, d || e)
	fmt.Println(!d, !e)

	// 位运算
	a = 60
	b = 13

	fmt.Println(a&b, a|b, a^b, a<<1, b<<2)

}
