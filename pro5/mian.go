package main

import (
	"fmt"
)

const (
	// iota 为const变量的数量计数，每增加一个就+1 直到下一个const关键字出现
	a = iota
	b
	c
	d
	e = "haha"
	f = "1s"
	g = iota // 恢复计数
	h
)

const (
	a1 = iota
	a2
	a3
	a4
)

func main() {
	fmt.Println("pro4")
	fmt.Println(a, b, c, d, e, f, g, h) // 0 1 2 3 haha 1s 6 7
	fmt.Println(a1, a2, a3, a4)         // 0 1 2 3
}
