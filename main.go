package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	fmt.Println("server running")
	err := http.ListenAndServe(":31800", nil)
	if err != nil {
		fmt.Println("server run err %v", err)
	}
}
