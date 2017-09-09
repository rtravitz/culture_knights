package router

import (
	"log"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rtravitz/culture_knights/club"
	"github.com/rtravitz/culture_knights/db"
)

func NewRouter() *chi.Mux {
	database, err := db.New(os.Getenv("CULTURE_DB"))
	if err != nil {
		log.Fatal("Could not connect to database: ", err.Error())
	}
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	initializeRoutes(r, database)

	return r
}

func initializeRoutes(r *chi.Mux, db *db.DB) {
	bookService := club.BookService{Key: os.Getenv("BOOKS_KEY"), Base: "https://www.googleapis.com/books/v1/"}
	bookEnv := &club.Env{DB: db, Service: bookService}

	r.Route("/users", func(r chi.Router) {
		r.Get("/", club.GetUsersHandler(db))
		r.Post("/", club.CreateUser(db))
		r.Get("/{id:[0-9]+}", club.GetUser(db))
		r.Put("/{id:[0-9]+}", club.UpdateUser(db))
		r.Delete("/{id:[0-9]+}", club.DeleteUser(db))
	})

	r.Route("/books", func(r chi.Router) {
		r.Post("/", bookEnv.CreateBook)
		r.Get("/", bookEnv.GetBooks)
	})

	r.Route("/meetings", func(r chi.Router) {
		r.Post("/", club.CreateMeeting(db))
		r.Get("/", club.GetMeetings(db))
	})
}
