package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	/books - get all books
	/books/id - get book by id
	/books - POST - create new book
*/

type Books struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

var MyBooks = []Books{
	Books{
		Id:     "1",
		Author: "Достоевский",
		Title:  "Преступление и наказание",
	},
	Books{
		Id:     "2",
		Author: "Толстой",
		Title:  "Война и мир",
	},
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MyBooks)
}
func createBook(w http.ResponseWriter, r *http.Request) {

}
func getBookById(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", getBookById).Methods("GET")

	http.ListenAndServe(":8080", r)
}
