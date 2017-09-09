package club

import "github.com/rtravitz/culture_knights/db"

type Book struct {
	ID            int     `json:"id,omitempty"`
	Title         string  `json:"title,omitempty"`
	Author        string  `json:"author,omitempty"`
	PublishedDate string  `json:"published_date,omitempty"`
	PageCount     int     `json:"page_count,omitempty"`
	AverageRating float64 `json:"average_rating,omitempty"`
	Thumbnail     string  `json:"thumbnail,omitempty"`
	Description   string  `json:"description,omitempty"`
	Link          string  `json:"link,omitempty"`
}

func (b *Book) Create(db *db.DB) error {
	err := db.QueryRow("INSERT INTO books(title, author, publishedDate, pageCount, averageRating, thumbnail, description, link) "+
		"Values($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		b.Title, b.Author, b.PublishedDate, b.PageCount, b.AverageRating, b.Thumbnail, b.Description, b.Link).Scan(&b.ID)
	if err != nil {
		return err
	}

	return nil
}

func (b *Book) Get(db *db.DB) error {
	return db.QueryRow("SELECT * FROM books WHERE id=$1", b.ID).Scan(
		&b.ID, &b.Title, &b.Author, &b.PublishedDate,
		&b.PageCount, &b.AverageRating, &b.Thumbnail,
		&b.Description, &b.Link,
	)
}

func All(db *db.DB) ([]Book, error) {
	rows, err := db.Query("SELECT * FROM books")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []Book{}

	for rows.Next() {
		var b Book

		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.PublishedDate, &b.PageCount,
			&b.AverageRating, &b.Thumbnail, &b.Description, &b.Link); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}
