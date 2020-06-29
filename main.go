package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     string
	Name   string
	Author *Author
}

type Author struct {
	Firstname string
	Lastname  string
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	//book.Id = strconv.Itoa(rand.Intn(100))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	//dummy book data
	books = append(books, Book{Id: "1", Name: "2 states", Author: &Author{Firstname: "chetan", Lastname: "bhagat"}})
	books = append(books, Book{Id: "2", Name: "Pride And PreJudice", Author: &Author{Firstname: "Jane", Lastname: "Austen"}})
	books = append(books, Book{Id: "3", Name: "Romeo And Juliet", Author: &Author{Firstname: "William", Lastname: "Shakespeare"}})

	r.HandleFunc("/app/books", getBooks).Methods("GET")
	r.HandleFunc("/app/books/{id}", getBookById).Methods("GET")
	r.HandleFunc("/app/books", createBook).Methods("POST")
	r.HandleFunc("/app/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/app/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
