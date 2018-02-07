package main

import (
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("login.html")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
