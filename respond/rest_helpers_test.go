package respond_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/rtravitz/culture_knights/respond"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type ExamplePayload struct {
	ExampleString string
	ExampleInt    int
}

var _ = Describe("RestHelpers", func() {
	Context("WithJSON", func() {
		It("Writes a json payload, response code, and headers", func() {
			payload := ExamplePayload{"You know nothing", 49}
			recorder := httptest.NewRecorder()
			WithJSON(recorder, http.StatusOK, payload)

			var result ExamplePayload
			json.NewDecoder(recorder.Body).Decode(&result)

			Expect(recorder.Code).To(Equal(200))
			Expect(recorder.HeaderMap["Content-Type"][0]).To(Equal("application/json"))
			Expect(result.ExampleString).To(Equal("You know nothing"))
			Expect(result.ExampleInt).To(Equal(49))
		})
	})

	Context("WithError", func() {
		It("Writes a status code and a json error message", func() {
			recorder := httptest.NewRecorder()
			WithError(recorder, http.StatusInternalServerError, "Chaos is a ladder.")

			var result map[string]string
			json.NewDecoder(recorder.Body).Decode(&result)

			Expect(recorder.Code).To(Equal(500))
			Expect(recorder.HeaderMap["Content-Type"][0]).To(Equal("application/json"))
			Expect(result["error"]).To(Equal("Chaos is a ladder."))
		})
	})
})
