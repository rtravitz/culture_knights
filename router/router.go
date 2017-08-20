package router

import (
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/rtravitz/culture_knights/books"
	"github.com/rtravitz/culture_knights/db"
	"github.com/rtravitz/culture_knights/users"
)

func NewRouter() *chi.Mux {
	database, err := db.New(os.Getenv("CULTURE_DB"))
	if err != nil {
		log.Fatal("Could not connect to database: ", err.Error())
	}
	r := chi.NewRouter()
	initializeRoutes(r, database)

	return r
}

func initializeRoutes(r *chi.Mux, db *db.DB) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", users.GetUsersHandler(db))
		r.Post("/", users.CreateUser(db))
		r.Get("/{id:[0-9]+}", users.GetUser(db))
		r.Put("/{id:[0-9]+}", users.UpdateUser(db))
		r.Delete("/{id:[0-9]+}", users.DeleteUser(db))
	})

	r.Route("/books", func(r chi.Router) {
		r.Post("/", books.CreateBook(db))
		r.Get("/", books.GetBooks(db))
	})
}
