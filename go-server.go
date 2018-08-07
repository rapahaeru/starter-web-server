package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func showPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                                  // parse arguments, you have to call this by yourself
	fmt.Println(r)                                 // shows entire request
	fmt.Println("form info server side: ", r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["algumacoisa"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

func main() {
	http.HandleFunc("/", showPage)           // seta o router
	err := http.ListenAndServe(":9000", nil) // seta porta
	if err != nil {
		log.Fatal("ListenServe: ", err)
	}

}
