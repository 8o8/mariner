package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"io"
)

func main() {
	fmt.Println("Mariner webd starting...")

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)

	r.HandleFunc("/voyages", indexHandler)


	n := negroni.Classic()
	n.UseHandler(r)
	n.Run()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "oi")
}

func voyageHandler(w http.ResponseWriter, r *http.Request) {

}