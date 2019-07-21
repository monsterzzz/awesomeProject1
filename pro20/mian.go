package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 简易web服务

	http.HandleFunc("/", indexPage)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello url_path = %s ", r.URL.Path)
}
