package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const port = ":8080"

type credentials struct {
	UsernameOrEmail string
	Password        string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("POST")

	// Solves Cross Origin Access Issue
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})

	srv := &http.Server{
		Handler: c.Handler(r),
		Addr:    port,
	}

	log.Println("Started Server at port ", port)

	log.Fatal(srv.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi")
	fmt.Fprintf(w, "Hello World!")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var creds credentials
	err := decoder.Decode(&creds)

	if err != nil {
		panic(err)
	}
	log.Println(creds.UsernameOrEmail)
}
