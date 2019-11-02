package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	rtr := mux.NewRouter()

	rtr.HandleFunc("/{topic}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Write([]byte(vars["topic"]))
	})

	rtr.HandleFunc("/data/{name}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is Sample Module"))
	})

	http.Handle("/", rtr)

	http.ListenAndServe(":4000", nil)
}
