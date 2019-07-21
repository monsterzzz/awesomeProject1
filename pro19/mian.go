package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 并发网络请求
	var l = []string{"http://tieba.baidu.com/p/6154529755", "http://tieba.baidu.com/p/6154529755", "http://tieba.baidu.com/p/6154529755"}
	c := make(chan string)
	for _, v := range l {
		go feachAll(v, c)
	}
	for range l {
		fmt.Printf("%s", <-c)
	}
	//for a:= 0 ; a<len(l); a++{
	//	fmt.Printf("%s\n",<-c)
	//}

}

func feachAll(url string, c chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprint(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c <- fmt.Sprint(err)
	}
	resp.Body.Close()
	//fmt.Printf("%s",b)
	//c <- fmt.Sprint(b)
	//fmt.Printf("%s",b)
	c <- fmt.Sprintf("%s", b)

}
