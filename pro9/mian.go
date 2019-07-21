package main

import (
	"fmt"
	"strconv"
)

func firstGoFunc() {
	fmt.Println("hello,go_func")
}

func secondFun(num1, num2 int) int { // 返回值类型需要声明
	num := num1 + num2
	return num
}

func threeFun(x, y string) string {
	return x + y
}

func fourFun(x string, y int) string {
	i := strconv.Itoa(y)
	return x + i
}

func fiveFun(x, y string, z int) (string, string) { // 多返回值
	return x + strconv.Itoa(z), y + strconv.Itoa(2)
}

func main() {
	fmt.Println("pro9")
	firstGoFunc()
	fmt.Println(secondFun(1, 2))
	fmt.Println(threeFun("hello", "world"))
	for a := 0; a < 10; a++ {
		fmt.Println(fourFun("hello", a))
	}
	for a := 0; a < 10; a++ {
		fmt.Println(fiveFun("hello", "world", a))
	}
}
