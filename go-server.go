package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

func main() {

	Router := chi.NewRouter()
	Router.Get("/", showPage)

	http.ListenAndServe(":9090", Router)
}

type Emails struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Img                  string `json:"img"`
	DateDelivery         string `json:"date_delivery"`
	DateDeliveryFormated string `json:"date_delivery_formated"`
}

func showPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself
	// fmt.Println(r)                                 // shows entire request
	fmt.Println("form info server side: ", r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["algumacoisa"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // send data to client side

	json, _ := responseJSON()

	w.Write(json)

	fmt.Println("json : ", string(json))
}

func responseJSON() ([]byte, error) {
	//https://tour.golang.org/moretypes/20
	data := []Emails{
		{1, "Cael", "Descrição do cael", "/assets/images/hyoga.gif", "2018-06-25", "25/06/2018"},
		{2, "ueda", "Descrição do Ueda", "/assets/images/seiya.jpg", "2018-06-20", "20/06/2018"},
		{2, "celestino", "Descrição do Celestino", "/assets/images/shun.png", "2018-05-24", "24/05/2018"},
		{2, "vasco", "Descrição do Vasco", "/assets/images/vegeta.jpg", "2018-03-15", "15/03/2018"},
	}

	return json.Marshal(data)
}
