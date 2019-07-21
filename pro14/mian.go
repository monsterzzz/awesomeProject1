package main

import (
	"fmt"
	"time"
)

func say(s string) {

	for i := 0; i <= 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Duration(1) * time.Second)
	}
}

func main() {
	fmt.Println("pro14")
	go say("hello")
	say("world")

}
