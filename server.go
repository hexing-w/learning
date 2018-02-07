package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Name(w http.ResponseWriter, request *http.Request) {
	t := request.URL.Path[1:]
	//fmt.Fprintf(w, strings.ToUpper(t))
	//fmt.Fprintf(w, request.URL.Path[1:]+strings.ToUpper(t))
}

func main() {
	http.HandleFunc("/", Name)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
