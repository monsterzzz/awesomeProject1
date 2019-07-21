package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	// 网络请求

	urlLi := make([]string, 3)

	urlLi = append(urlLi, "https://www.baidu.com")

	for _, v := range urlLi {
		// http://
		if strings.HasPrefix(v, "http://") {
			v = "http://" + v
		}
		if v == "" {
			continue
		}
		resp, err := http.Get(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		//dst ,err:= io.Copy(os.Stdout,resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", resp.Status)
		fmt.Printf("%s\n", b)
	}

}
