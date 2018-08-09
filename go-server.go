package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rapahaeru/starter-web-server/api"
)

func main() {

	Router := chi.NewRouter()
	Router.Get("/", handler)

	http.ListenAndServe(":9090", Router)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, _ := api.GetEmailList()
	w.Write(json)

	fmt.Println("json : ", string(json))
}
