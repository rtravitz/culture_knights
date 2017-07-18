package books

import (
	"database/sql"
	"net/http"
	"os"
)

func GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := All(db)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, books)
	}
}

func CreateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.FormValue("q")
		if q == "" {
			respondWithError(w, http.StatusBadRequest, "Please send a query")
		}
		bookService := NewService(os.Getenv("BOOKS_KEY"), "https://www.googleapis.com/books/v1/")

		book, err := bookService.FindBook(q)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := book.Create(db); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, book)
	}
}
