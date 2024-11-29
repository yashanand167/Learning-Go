package main

import (
	"log"
	"net/http"
	"my-app/controllers"
)

func main(){
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request){
		w.Write([] byte ("Hello World"))
	})
	http.HandleFunc("/api/books", controllers.GetBooks)
	http.HandleFunc("/api/books/add", controllers.AddBook)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}