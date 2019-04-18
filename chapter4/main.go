package main

import (
	"fmt"
	"log"
	"net/http"

	booksHandler "github.com/kind84/distributed-go/chapter4/books-handler"
)

func main() {
	books := booksHandler.GetBooks()
	log.Println(fmt.Sprintf("%+v", books))

	actionCh := make(chan booksHandler.Action)

	go booksHandler.StartBooksManager(books, actionCh)

	http.HandleFunc("/api/books/", booksHandler.MakeHandler(booksHandler.BookHandler, "/api/books/", actionCh))

	log.Println("Starting server at port 8080...")
	http.ListenAndServe(":8080", nil)
}
