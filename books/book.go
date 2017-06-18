package books

import "database/sql"

type Book struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	PublishedDate string  `json:"published_date"`
	PageCount     int     `json:"page_count"`
	AverageRating float64 `json:"average_rating"`
	Thumbnail     string  `json:"thumbnail"`
	Description   string  `json:"description"`
	Link          string  `json:"link"`
}

func (b *Book) CreateBook(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO books(title, author, publishedDate, pageCount, averageRating, thumbnail, description, link) "+
		"Values($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		b.Title, b.Author, b.PublishedDate, b.PageCount, b.AverageRating, b.Thumbnail, b.Description, b.Link).Scan(&b.ID)
	if err != nil {
		return err
	}

	return nil
}
