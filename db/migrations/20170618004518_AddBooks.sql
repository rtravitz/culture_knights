
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE books
(
  id Serial,
  title TEXT,
  author TEXT,
  publishedDate TEXT,
  pageCount INT,
  averageRating REAL,
  thumbnail TEXT,
  description TEXT,
  link TEXT
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE books;

