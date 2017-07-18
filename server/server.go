package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/rtravitz/culture_knights/books"
	"github.com/rtravitz/culture_knights/users"
)

func StartServer() {
	db, err := OpenDB(os.Getenv("CULTURE_DB"))
	if err != nil {
		log.Fatal("Could not connect to database: ", err.Error())
	}
	r := mux.NewRouter()
	initializeRoutes(r, db)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}

func OpenDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initializeRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/users/{id:[0-9]+}", users.GetUser(db)).Methods("GET")
	r.HandleFunc("/users", users.GetUsersHandler(db)).Methods("GET")
	r.HandleFunc("/users", users.CreateUser(db)).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", users.DeleteUser(db)).Methods("DELETE")
	r.HandleFunc("/users/{id:[0-9]+}", users.UpdateUser(db)).Methods("PUT")
	r.HandleFunc("/books", books.CreateBook(db)).Methods("POST")
	r.HandleFunc("/books", books.GetBooks(db)).Methods("GET")
}
