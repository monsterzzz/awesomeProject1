package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("pro3")
	// 常量
	const s1 string = "hello"
	const s2 = "world"
	const s3, s4, s5, s6 = "1", "2", "3", 4
	var addS1, addS2 string // 没有定义值的时候必须声明类型
	addS1 = s1 + " " + s2
	addS2 = s3 + s4 + s5
	fmt.Println(addS1)
	fmt.Println(addS2)
	//fmt.Println(addS1 + s6) error 字符串和整形相加
	fmt.Println(s6)

	var ss1 = "abc"
	var a = len(ss1)
	var b = unsafe.Sizeof(ss1) // 第一个域是指向该字符串的指针，第二个域是字符串的长度，每个域占8个字节，但是并不包含指针指向的字符串的内容
	// 32位系统就是  4 + 4 = 8
	fmt.Println(ss1, a, b) // abc 3 16

}
