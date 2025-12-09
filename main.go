package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*
	/books - get all books
	/books/id - get book by id
	/books - POST - create new book
*/

type Books struct {
	Id     string
	Author string
	Title  string
}

func getBooks(w http.ResponseWriter, r *http.Request) {

}
func createBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")

	http.ListenAndServe(":8080", r)
}
