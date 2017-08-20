package books

import (
	"net/http"

	"github.com/rtravitz/culture_knights/db"
	"github.com/rtravitz/culture_knights/respond"
)

type Env struct {
	DB *db.DB
	Service
}

func (env *Env) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := All(env.DB)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respond.WithJSON(w, http.StatusOK, books)
}

func (env *Env) CreateBook(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	if q == "" {
		respond.WithError(w, http.StatusBadRequest, "Please send a query")
	}

	book, err := env.Service.FindBook(q)
	if err != nil {
		respond.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := book.Create(env.DB); err != nil {
		respond.WithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respond.WithJSON(w, http.StatusOK, book)
}
