package data

type Book struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
type BOOKS []*Book

func GetBooks() BOOKS {
	return bookslist
}
func NewBook() *Book {
	return &Book{}
}
func AddBook(newBook *Book) {
	newBook.ID = bookslist[len(bookslist)-1].ID + 1
	bookslist = append(bookslist, newBook)
}
func DeleteBook(idInt int64) {
	for i, book := range bookslist {
		if book.ID == idInt {
			bookslist = append(bookslist[:i], bookslist[i+1:]...)
			break
		}
	}
}
func Update() {
	for i, book := range bookslist {
		book.ID = int64(i) + 1
	}
}

var bookslist = BOOKS{
	{
		ID:     1,
		Title:  "test",
		Author: "test2",
		Price:  14.99,
	},

	{
		ID:     2,
		Title:  "test3",
		Author: "test4",
		Price:  69.99,
	},
}
