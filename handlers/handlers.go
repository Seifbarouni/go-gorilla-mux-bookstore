package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"projects/Bookstore/data"
	"strconv"

	"github.com/gorilla/mux"
)

type BookHanlder struct{}

func NewHandler() *BookHanlder {
	return &BookHanlder{}
}

var tpl *template.Template

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	tpl = template.Must(template.ParseGlob("assets/*.html"))
}

func (bh *BookHanlder) Index(rw http.ResponseWriter, r *http.Request) {
	http.Redirect(rw, r, "/books/show", http.StatusFound)
}

func (bh *BookHanlder) Books(rw http.ResponseWriter, r *http.Request) {
	log.Println("Handling GET Request : books")
	book := data.GetBooks()

	err := tpl.ExecuteTemplate(rw, "books.html", book)
	handleErr(err)
}
func (bh *BookHanlder) Add(rw http.ResponseWriter, r *http.Request) {
	log.Println("Handling POST Request : add")

	err := tpl.ExecuteTemplate(rw, "AddForm.html", nil)
	handleErr(err)

}
func (bh *BookHanlder) AddProcess(rw http.ResponseWriter, r *http.Request) {
	newBook := data.NewBook()
	var err error
	newBook.Title = r.FormValue("title")
	newBook.Author = r.FormValue("author")
	newBook.Price, err = strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		rw.Write([]byte("the Price variable is a float"))
	}
	data.AddBook(newBook)
	log.Println("[SUCCESS]New Book Added.")
	http.Redirect(rw, r, "/books/show", http.StatusFound)

}
func (bh *BookHanlder) DeleteBook(rw http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(rw, "DeleteBook.html", nil)
	handleErr(err)

}
func (bh *BookHanlder) DeleteProcess(rw http.ResponseWriter, r *http.Request) {
	id := r.FormValue("ID")
	idInt, err := strconv.ParseInt(id, 10, 64)
	handleErr(err)
	data.DeleteBook(idInt)
	log.Println("[SUCCESS]Book Deleted.")
	data.Update()
	http.Redirect(rw,r,"/books/show",http.StatusFound)

}
func (bh *BookHanlder) Update(rw http.ResponseWriter, r *http.Request) {
	Books := data.GetBooks()
	err := tpl.ExecuteTemplate(rw, "UpdateBook.html", Books)
	handleErr(err)
}
func (bh *BookHanlder) UpdateProcess(rw http.ResponseWriter, r *http.Request) {
	Books := data.GetBooks()
	id, err := strconv.ParseInt(r.FormValue("ID"), 10, 64)
	handleErr(err)
	for i, book := range Books {
		if book.ID == id {
			if r.FormValue("author") != "" {
				book.Author = r.FormValue("author")
			}
			if r.FormValue("title") != "" {
				book.Title = r.FormValue("title")
			}
			if r.FormValue("price") != "" {
				book.Price, err = strconv.ParseFloat(r.FormValue("price"), 64)
				handleErr(err)
			}

			break
		}
		if i >= len(Books) {
			rw.Write([]byte("Cannot find the Book.."))
			return
		}
	}
	data.Update()
	http.Redirect(rw, r, "/books/show", http.StatusFound)

}
func (bh *BookHanlder) ShowAPI(rw http.ResponseWriter, r *http.Request) {
	log.Println("Handling API Get Request..")
	Books := data.GetBooks()
	rw.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(rw)
	err := e.Encode(Books)
	handleErr(err)
}
func (bh *BookHanlder) AddBookAPI(rw http.ResponseWriter, r *http.Request) {
	log.Println("Handling API Post Request..")

	body := r.Body
	book := data.NewBook()
	d := json.NewDecoder(body)
	err := d.Decode(book)
	handleErr(err)
	data.AddBook(book)

}
func (bh *BookHanlder) DeleteBookAPI(rw http.ResponseWriter, r *http.Request) {
	log.Println("Handling API Delete Request..")

	vars := mux.Vars(r)
	Books := data.GetBooks()
	idString := vars["id"]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		rw.WriteHeader(400)
		rw.Write([]byte("Could not convert ID to int64"))
		return
	}
	if id >= int64(len(Books)) {
		rw.WriteHeader(404)
		rw.Write([]byte("No post found with specified ID"))
		return
	}

	for i, book := range Books {
		if book.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			break
		}
	}

}
