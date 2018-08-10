package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/rapahaeru/starter-web-server/api"
)

func main() {

	Router := chi.NewRouter()
	Router.Get("/", handler)

	http.ListenAndServe(":9090", Router)
}

func handler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")

	for index, value := range r.Form {

		if index == "token" {
			token := strings.Join(value, "")
			fmt.Printf("retorno: %v", checkToken(token))

			if checkToken(token) {
				w.WriteHeader(http.StatusOK)
				json, _ := api.GetEmailList()
				w.Write(json)
			} else {
				fmt.Fprintf(w, "Faltou o token do Miguelito")
			}
		}
	}

}

func checkToken(token string) bool {

	tokenList := api.GetToken()

	for _, value := range tokenList {
		if value.Token == token {
			return true
		}
	}

	return false
}
