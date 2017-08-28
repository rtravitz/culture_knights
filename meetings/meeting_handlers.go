package meetings

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rtravitz/culture_knights/db"
	"github.com/rtravitz/culture_knights/respond"
)

func CreateMeeting(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m Meeting
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			log.Println("CreateMeeting Handler Err: ", err.Error())
			respond.WithError(w, http.StatusInternalServerError, "Something went wrong!")
			return
		}
		defer r.Body.Close()

		if err := m.Create(db); err != nil {
			log.Println("CreateMeeting Handler Err: ", err.Error())
			respond.WithError(w, http.StatusInternalServerError, "Something went wrong!")
			return
		}

		m.Book.Get(db)
		m.JesusChooser.Get(db)
		m.BookChooser.Get(db)
		m.GoldenRoosterRecipient.Get(db)
		m.LeadBallsRecipient.Get(db)

		respond.WithJSON(w, http.StatusCreated, m)
	}
}

func GetMeetings(db *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		meetings, err := GetAll(db)
		if err != nil {
			log.Println("GetMeetings Handler Err: ", err.Error())
			respond.WithError(w, http.StatusInternalServerError, "Something went wrong!")
		}

		respond.WithJSON(w, http.StatusOK, meetings)
	}
}
