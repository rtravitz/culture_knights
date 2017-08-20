package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rtravitz/culture_knights/db"
	"github.com/rtravitz/culture_knights/respond"
)

func GetUsersHandler(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := GetAll(db)
		if err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respond.WithJSON(w, http.StatusOK, users)
	}
}

func GetUser(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(urlID)
		if err != nil {
			respond.WithError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}

		user := User{ID: id}
		if err := user.Get(db); err != nil {
			switch err {
			case sql.ErrNoRows:
				respond.WithError(w, http.StatusNotFound, "User not found")
			default:
				respond.WithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		respond.WithJSON(w, http.StatusOK, user)
	}
}

func CreateUser(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer r.Body.Close()

		if err := user.Create(db); err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respond.WithJSON(w, http.StatusCreated, user)
	}
}

func UpdateUser(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User

		urlID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(urlID)
		if err != nil {
			respond.WithError(w, http.StatusBadRequest, "Invalid User ID")
			return
		}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		defer r.Body.Close()
		user.ID = id

		if err = user.Update(db); err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respond.WithJSON(w, http.StatusOK, user)
	}
}

func DeleteUser(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(urlID)
		if err != nil {
			respond.WithError(w, http.StatusBadRequest, "Invalid User ID")
			return
		}

		user := User{ID: id}
		if err := user.Delete(db); err != nil {
			respond.WithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respond.WithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	}
}
