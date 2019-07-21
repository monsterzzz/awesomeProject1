package main

import (
	"fmt"
	"os"
)

func main() {
	// 命令行参数
	// os.Args

	fmt.Println("ll")
	fmt.Println(os.Args)
	for a := 0; a < len(os.Args); a++ {
		fmt.Printf("arg%d -> %s\n", a+1, os.Args[a])
	}

}
