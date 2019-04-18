package bookshandler

import "net/http"

func actOnGET(books map[string]BookResource, act Action) {
	status := http.StatusNotFound
	bookResult := []BookResource{}

	if act.ID == "" {
		status = http.StatusOK
		for _, book := range books {
			bookResult = append(bookResult, book)
		}
	} else if book, exists := books[act.ID]; exists {
		status = http.StatusOK
		bookResult = []BookResource{book}
	}

	act.RetChan <- response{
		StatusCode: status,
		Books:      bookResult,
	}
}

func actOnDELETE(books map[string]BookResource, act Action) {
	book, exists := books[act.ID]
	delete(books, act.ID)

	if !exists {
		book = BookResource{}
	}

	act.RetChan <- response{
		StatusCode: http.StatusOK,
		Books:      []BookResource{book},
	}
}

func actOnPUT(books map[string]BookResource, act Action) {
	status := http.StatusNotFound
	bookResult := []BookResource{}

	if book, exists := books[act.ID]; exists {
		book.Link = act.Payload.Link
		book.Title = act.Payload.Title
		books[act.ID] = book

		status = http.StatusOK
		bookResult = []BookResource{books[act.ID]}
	}

	act.RetChan <- response{
		StatusCode: status,
		Books:      bookResult,
	}
}

func actOnPOST(books map[string]BookResource, act Action, newID string) {
	books[newID] = BookResource{
		ID:    newID,
		Link:  act.Payload.Link,
		Title: act.Payload.Title,
	}

	act.RetChan <- response{
		StatusCode: http.StatusCreated,
		Books:      []BookResource{books[newID]},
	}
}
