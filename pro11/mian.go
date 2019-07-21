package main

import (
	"fmt"
)

type people struct {
	name string
	sex  int
	age  int
}

func printPeople(p people) {
	fmt.Printf("people name : %s\n", p.name)
	fmt.Printf("people sex : %d\n", p.sex)
	fmt.Printf("people age : %d\n", p.age)
}

func (p people) say(s string) {
	fmt.Printf("%s say: %s", p.name, s)
}

func main() {
	fmt.Println("pro11")

	var p1 people

	p1.name = "mm"
	p1.age = 10
	p1.sex = 0

	printPeople(p1)
	p1.say("hello")

}
