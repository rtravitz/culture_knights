package users_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi"
	. "github.com/rtravitz/culture_knights/users"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var recorder *httptest.ResponseRecorder

var _ = Describe("UserHandlers", func() {
	BeforeEach(func() {
		populateUsers()
		recorder = httptest.NewRecorder()
	})

	AfterEach(func() {
		testDB.Exec(`TRUNCATE users`)
		testDB.Exec(`ALTER SEQUENCE users_id_seq RESTART WITH 1`)
	})

	Context("GetUsersHandler", func() {
		It("Returns a list of all users", func() {
			req, _ := http.NewRequest("GET", "/users", nil)
			handler := GetUsersHandler(testDB)
			handler.ServeHTTP(recorder, req)

			var result []User
			json.NewDecoder(recorder.Body).Decode(&result)

			Expect(recorder.Code).To(Equal(200))
			Expect(result[0].Name).To(Equal("Jon"))
			Expect(result[2].Name).To(Equal("Arya"))
		})
	})

	Context("GetUser", func() {
		It("Returns a matching user if the ID exists", func() {
			req, _ := http.NewRequest("GET", "/users/2", nil)
			router().ServeHTTP(recorder, req)

			var result User
			json.NewDecoder(recorder.Body).Decode(&result)

			Expect(recorder.Code).To(Equal(200))
			Expect(result.Name).To(Equal("Jorah"))
			Expect(result.ID).To(Equal(2))
		})

		It("Returns a not found message if the user can't be located", func() {
			req, _ := http.NewRequest("GET", "/users/5", nil)
			router().ServeHTTP(recorder, req)

			var result map[string]string
			json.NewDecoder(recorder.Body).Decode(&result)

			Expect(recorder.Code).To(Equal(404))
			Expect(result["error"]).To(Equal("User not found"))
		})
	})
})

func populateUsers() {
	users := []User{User{Name: "Jon"}, User{Name: "Jorah"}, User{Name: "Arya"}}
	for _, user := range users {
		user.Create(testDB)
	}
}

func router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/users/{id:[0-9]+}", GetUser(testDB))
	return r
}