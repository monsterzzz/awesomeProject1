package main

import (
	"fmt"
)

func main() {
	fmt.Println("pro12")

	// 切片
	s1 := []int{1, 2, 3}
	s2 := make([]int, 3)

	fmt.Println(s1)
	s1 = append(s1, 4)
	fmt.Println(s1)

	fmt.Println(s2)
	s2 = append(s2, 1)
	fmt.Println(s2)

	s3 := s1[:3]
	fmt.Println(s3)
	fmt.Println(len(s3), cap(s3), s3)
	s3 = append(s3, 66)
	fmt.Println(len(s3), cap(s3), s3)
	s3 = append(s3, 66)
	fmt.Println(len(s3), cap(s3), s3)
	s3 = append(s3, 66)
	fmt.Println(len(s3), cap(s3), s3)
	s3 = append(s3, 66)
	fmt.Println(len(s3), cap(s3), s3) // cap 容量 len 长度

}
