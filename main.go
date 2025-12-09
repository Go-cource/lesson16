package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")
	var newBook Books
	// err := json.NewDecoder(r.Body).Decode(&newBook)
	// if err != nil {
	// 	log.Println("Error while decoding: ", err)
	// 	fmt.Fprint(w, "504 Internal Error")
	// 	return
	// }
	newBook.Author = r.FormValue("Author")
	newBook.Title = r.FormValue("Title")

	newBook.Id = strconv.Itoa(len(MyBooks) + 1)
	MyBooks = append(MyBooks, newBook)
	fmt.Fprint(w, "200 OK")
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for _, item := range MyBooks {
		if item.Id == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	fmt.Fprint(w, "No such Book, Sorry!")

}

func createBookInterface(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/create_book", createBookInterface).Methods("GET")

	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", getBookById).Methods("GET")

	log.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
