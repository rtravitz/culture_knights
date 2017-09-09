package club

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Service interface {
	FindBook(string) (*Book, error)
}

type BookService struct {
	Key  string
	Base string
}

type VolumeInfo struct {
	Title         string   `json:"title"`
	Authors       []string `json:"authors"`
	Publisher     string   `json:"publisher"`
	PublishedDate string   `json:"publishedDate"`
	Description   string   `json:"description"`
	PageCount     int      `json:"pageCount"`
	AverageRating float64  `json:"averageRating"`
	ImageLinks    struct {
		Thumbnail string `json:"thumbnail"`
	} `json:"imageLinks"`
	CanonicalVolumeLink string `json:"canonicalVolumeLink"`
}

type bookResponse struct {
	Items []struct {
		VolumeInfo `json:"volumeInfo"`
	} `json:"items"`
}

func (service BookService) FindBook(query string) (*Book, error) {
	u, err := url.Parse(service.Base + "volumes")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("q", query)
	q.Set("key", service.Key)
	u.RawQuery = q.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	br := new(bookResponse)
	json.NewDecoder(resp.Body).Decode(&br)
	vol := br.Items[0].VolumeInfo
	book := Book{
		Title:         vol.Title,
		Author:        vol.Authors[0],
		PublishedDate: vol.PublishedDate,
		PageCount:     vol.PageCount,
		AverageRating: vol.AverageRating,
		Thumbnail:     vol.ImageLinks.Thumbnail,
		Description:   vol.Description,
		Link:          vol.CanonicalVolumeLink,
	}

	return &book, nil
}
