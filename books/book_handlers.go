package books

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/rtravitz/culture_knights/respond"
)

func GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := All(db)
		if err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respond.WithJSON(w, http.StatusOK, books)
	}
}

func CreateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.FormValue("q")
		if q == "" {
			respond.WithError(w, http.StatusBadRequest, "Please send a query")
		}
		bookService := NewService(os.Getenv("BOOKS_KEY"), "https://www.googleapis.com/books/v1/")

		book, err := bookService.FindBook(q)
		if err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := book.Create(db); err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respond.WithJSON(w, http.StatusOK, book)
	}
}
