package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rtravitz/culture_knights/respond"
)

func GetUsersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, _ := strconv.Atoi(r.FormValue("count"))
		start, _ := strconv.Atoi(r.FormValue("start"))

		if count > 10 || count < 1 {
			count = 10
		}
		if start < 0 {
			start = 0
		}

		users, err := getUsers(db, start, count)
		if err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respond.WithJSON(w, http.StatusOK, users)
	}
}

func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(urlID)
		if err != nil {
			respond.WithError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}

		u := User{ID: id}
		if err := u.getUser(db); err != nil {
			switch err {
			case sql.ErrNoRows:
				respond.WithError(w, http.StatusNotFound, "User not found")
			default:
				respond.WithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		respond.WithJSON(w, http.StatusOK, u)
	}
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&u); err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer r.Body.Close()

		if err := u.createUser(db); err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respond.WithJSON(w, http.StatusCreated, u)
	}
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User

		urlID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(urlID)
		if err != nil {
			respond.WithError(w, http.StatusBadRequest, "Invalid User ID")
			return
		}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&u); err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer r.Body.Close()
		u.ID = id

		if err = u.updateUser(db); err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respond.WithJSON(w, http.StatusOK, u)
	}
}

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(urlID)
		if err != nil {
			respond.WithError(w, http.StatusBadRequest, "Invalid User ID")
			return
		}

		u := User{ID: id}
		if err := u.deleteUser(db); err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respond.WithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	}
}
