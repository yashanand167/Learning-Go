package controllers

import (
	"encoding/json"
	"net/http"
	"my-app/models"
)

// sample data
var books = []models.Book{
	{ID: 1, Title: "Go Programming", Author: "John Doe"},
	{ID: 2, Title: "Web Development", Author: "Jane Smith"},
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid book data", http.StatusBadRequest)
		return
	}

	newBook.ID = len(books) + 1 
	books = append(books, newBook)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBook)
}
