package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/bad", badPage).Methods("GET")
	router.HandleFunc("/name/{param}", nameParamPage).Methods("GET")
	router.HandleFunc("/data", dataPage).Methods("POST")
	router.HandleFunc("/headers", headersPage).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

// main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func badPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func nameParamPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("Hello, " + vars["param"] + "!"))
	w.WriteHeader(http.StatusOK)
}

func dataPage(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte("I got message:\n" + string(data)))
}

func headersPage(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	if a, ok := h["A"]; ok {
		if b, ok := h["B"]; ok {
			first, err := strconv.Atoi(a[0])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			second, err := strconv.Atoi(b[0])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			w.Header().Set("a+b", strconv.Itoa(first+second))
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}
