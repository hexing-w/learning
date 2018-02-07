package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello xingshuai")
}

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.tpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password", r.Form["password"])
		if len(r.Form["username"][0]) == 0 {
			fmt.Println("username can't null")
		}
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("listenAndServe:", err)
	}
}
