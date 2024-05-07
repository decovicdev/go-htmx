package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {

	fmt.Println("hello world")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("index.html"))

		tmpl.Execute(w, nil)

	})

	log.Fatal(
		http.ListenAndServe(":3000", nil),
	)

}
