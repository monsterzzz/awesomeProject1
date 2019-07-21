package main

import (
	"fmt"
)

func main() {
	fmt.Println("pro10")
	// 数组
	var li1 = [...]int{1, 1, 1, 1, 1, 1}
	var li2 = [6]int{1, 1, 1, 1, 1, 1}
	var li3 = li1
	fmt.Println(li1 == li2)
	fmt.Println(&li1 == &li2, &li1 == &li3)
	for i := 0; i < len(li1); i++ {
		fmt.Println(li1[i])
	}

	for i, v := range li1 {
		fmt.Printf("li[%d] = %d\n", i, v)
	}

}
