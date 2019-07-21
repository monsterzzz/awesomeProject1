package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 查找重复行

	//1
	//counts := make(map[string] int)
	//ipt := bufio.NewScanner(os.Stdin)
	//for ipt.Scan() {
	//	counts[ipt.Text()]++
	//}
	//
	//for line,n := range counts{
	//	fmt.Println(line,n)
	//}

	//2
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		sonMap := make(map[string]int)
		countLine(os.Stdin, sonMap)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println("error file! %s", file)
				continue
			}
			if counts[file] == nil {
				counts[file] = make(map[string]int)
			}
			countLine(f, counts[file])
		}
	}
	for file, v := range counts {
		for _, n := range v {
			if n >= 2 {
				fmt.Println(file)
				break
			}
		}
	}

}

func countLine(f *os.File, count map[string]int) {
	ipt := bufio.NewScanner(f)
	for ipt.Scan() {
		count[ipt.Text()]++
	}
}
