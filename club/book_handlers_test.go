package club_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/rtravitz/culture_knights/club"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	recorder *httptest.ResponseRecorder
	env      *Env
)

var _ = Describe("BookHandlers", func() {
	BeforeEach(func() {
		populateBooks()
		recorder = httptest.NewRecorder()
		mockService := &MockService{}
		env = &Env{DB: testDB, Service: mockService}
	})

	AfterEach(func() {
		testDB.Exec(`TRUNCATE books`)
		testDB.Exec(`ALTER SEQUENCE books_id_seq RESTART WITH 1`)
	})

	Context("GetBooks", func() {
		It("Returns all books", func() {
			req, _ := http.NewRequest("GET", "/books", nil)
			handler := http.HandlerFunc(env.GetBooks)
			handler.ServeHTTP(recorder, req)

			var result []Book
			json.NewDecoder(recorder.Body).Decode(&result)

			Expect(recorder.Code).To(Equal(200))
			Expect(result[0].Title).To(Equal("The Thirty-nine Steps"))
			Expect(result[1].PageCount).To(Equal(448))
		})
	})

	Context("CreateBook", func() {
		It("Creates and returns a book", func() {
			req, _ := http.NewRequest("POST", `/books?q="Devil in a Blue Dress`, nil)
			handler := http.HandlerFunc(env.CreateBook)
			handler.ServeHTTP(recorder, req)

			var result Book
			json.NewDecoder(recorder.Body).Decode(&result)

			Expect(recorder.Code).To(Equal(200))
			Expect(result.Title).To(Equal("Devil in a Blue Dress"))
			Expect(result.Author).To(Equal("Walter Mosley"))
		})

		It("Returns a status bad request if no query is sent", func() {
			req, _ := http.NewRequest("POST", `/books`, nil)
			handler := http.HandlerFunc(env.CreateBook)
			handler.ServeHTTP(recorder, req)

			var result map[string]string
			json.NewDecoder(recorder.Body).Decode(&result)

			Expect(recorder.Code).To(Equal(400))
			Expect(result["error"]).To(Equal("Please send a query"))
		})
	})
})

type MockService struct{}

func (mockService MockService) FindBook(query string) (*Book, error) {
	book := oneMockBook()
	return &book, nil
}

func populateBooks() {
	books := []Book{
		Book{Title: "The Thirty-nine Steps",
			Author:        "John Buchan",
			PublishedDate: "1915",
			PageCount:     231,
			AverageRating: 4.0,
			Thumbnail:     "Mock Thumbnail 1",
			Description:   "Mock Description 1",
			Link:          "Mock Link 1"},
		Book{Title: "The Martian",
			Author:        "Andy Weir",
			PublishedDate: "08-18-2015",
			PageCount:     448,
			AverageRating: 4.7,
			Thumbnail:     "Mock Thumbnail 2",
			Description:   "Mock Description 2",
			Link:          "Mock Link 2"},
	}

	for _, book := range books {
		book.Create(testDB)
	}
}

func oneMockBook() Book {
	return Book{Title: "Devil in a Blue Dress",
		Author:        "Walter Mosley",
		PublishedDate: "1990",
		PageCount:     219,
		AverageRating: 3.5,
		Thumbnail:     "Mock Thumbnail 3",
		Description:   "Mock Description 3",
		Link:          "Mock Link 3"}
}
