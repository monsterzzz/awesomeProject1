package main

import (
	"fmt"
)

func main() {
	fmt.Println("pro7")
	// 条件语句
	a := 20
	if a < 10 {
		fmt.Println(1)
	} else if a > 5 && a < 0 {
		fmt.Println("13")
	} else {
		fmt.Println("xxxx")
	}

	switch a { // 自动有break 如果找不到相应case 则会跳转到default 如无default则跳过
	case 20:
		fmt.Println("1")
		fallthrough // 强制执行后一个case
	case 30:
		fmt.Println("2")
	case 40:
		fmt.Println("3")
	case 50:
		fmt.Println("4")
	default:
		fmt.Println(a)
	}

	select { // wait
	default:
		fmt.Println("!")
	}

}
