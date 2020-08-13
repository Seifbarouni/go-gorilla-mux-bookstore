package main

import (
	"log"
	"net/http"
	"projects/BookStore/handlers"

	"github.com/gorilla/mux"
)

func main() {

	bh := mux.NewRouter()
	ph := handlers.NewHandler()

	shBookRouter := bh.Methods(http.MethodGet).Subrouter()
	shBookRouter.HandleFunc("/books/show", ph.Books)
	shBookRouter.HandleFunc("/", ph.Index)

	addBookRouter := bh.Methods(http.MethodGet).Subrouter()
	addBookRouter.HandleFunc("/books/add", ph.Add)

	addBookprocessRouter := bh.Methods(http.MethodPost).Subrouter()
	addBookprocessRouter.HandleFunc("/books/add/process", ph.AddProcess)

	deleteBookRouter := bh.Methods(http.MethodGet).Subrouter()
	deleteBookRouter.HandleFunc("/books/delete", ph.DeleteBook)

	deleteBookprocessRouter := bh.Methods(http.MethodGet).Subrouter()
	deleteBookprocessRouter.HandleFunc("/books/delete/process", ph.DeleteProcess)

	updateBook := bh.Methods(http.MethodGet).Subrouter()
	updateBook.HandleFunc("/books/update", ph.Update)

	updateBookprocess := bh.Methods(http.MethodGet).Subrouter()
	updateBookprocess.HandleFunc("/books/update/process", ph.UpdateProcess)

	apiRouter := bh.Methods(http.MethodGet).Subrouter()
	apiRouter.HandleFunc("/books/api", ph.ShowAPI)

	apiAddRouter := bh.Methods(http.MethodPost).Subrouter()
	apiAddRouter.HandleFunc("/books/api/add", ph.AddBookAPI)

	apiDeleteRouter := bh.Methods(http.MethodGet).Subrouter()
	apiDeleteRouter.HandleFunc("/books/api/delete/{id:[0-9]+}", ph.DeleteBookAPI)

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", bh))
}
