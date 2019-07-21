package main

import (
	"fmt"
)

func main() {
	fmt.Println("pro8")
	// for
	//1
	a := 5
	b := 20
	for a <= b {
		fmt.Println(a)
		a++
	}
	// 2
	for c := 0; c < b; c++ {
		fmt.Println("c value is ", c)
	}
	//3
	li := [6]int{0, 1, 2, 3, 6, 5}
	for idx, v := range li {
		fmt.Println(idx, v)
	}

	// 仿while
	d := true
	var p int
	for d {
		p++
		if p == 30 {
			d = false
		}
		fmt.Println(p, d)
	}
}
