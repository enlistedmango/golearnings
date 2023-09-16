package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmccallum-rbi/golearnings/crud/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Fatalf("Could not listen on port 4000  %v", err)
	}

}
