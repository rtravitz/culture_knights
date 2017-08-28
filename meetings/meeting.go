package meetings

import (
	"time"

	"github.com/rtravitz/culture_knights/books"
	"github.com/rtravitz/culture_knights/users"

	"github.com/rtravitz/culture_knights/db"
)

type Meeting struct {
	ID                     int        `json:"id,omitempty"`
	Book                   books.Book `json:"book,omitempty"`
	Jesus                  string     `json:"jesus,omitempty"`
	JesusChooser           users.User `json:"jesus_chooser,omitempty"`
	JesusExplanation       string     `json:"jesus_explanation,omitempty"`
	BookChooser            users.User `json:"book_chooser,omitempty"`
	GoldenRooster          string     `json:"golden_rooster,omitempty"`
	GoldenRoosterRecipient users.User `json:"golden_rooster_recipient,omitempty"`
	LeadBalls              string     `json:"lead_balls,omitempty"`
	LeadBallsRecipient     users.User `json:"lead_balls_recipient,omitempty"`
	DiscussionDate         time.Time  `json:"discussion_date,omitempty"`
}

func (m *Meeting) Create(db *db.DB) error {
	err := db.QueryRow(createQuery, m.Book.ID, m.Jesus, m.JesusChooser.ID, m.JesusExplanation,
		m.BookChooser.ID, m.GoldenRooster, m.GoldenRoosterRecipient.ID,
		m.LeadBalls, m.LeadBallsRecipient.ID, m.DiscussionDate).Scan(&m.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *Meeting) Get(db *db.DB) error {
	return db.QueryRow("SELECT * FROM meetings WHERE id=$1", m.ID).Scan(
		&m.Book.ID, &m.Jesus, &m.JesusChooser.ID, &m.JesusExplanation, &m.BookChooser.ID,
		&m.GoldenRooster, &m.GoldenRoosterRecipient.ID, &m.LeadBalls, &m.LeadBallsRecipient.ID,
		&m.DiscussionDate,
	)
}

func GetAll(db *db.DB) ([]Meeting, error) {
	rows, err := db.Query(getAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	meetings := []Meeting{}
	for rows.Next() {
		var m Meeting
		if err := rows.Scan(
			&m.ID, &m.Book.ID, &m.Jesus, &m.JesusChooser.ID, &m.JesusExplanation, &m.BookChooser.ID,
			&m.GoldenRooster, &m.GoldenRoosterRecipient.ID, &m.LeadBalls, &m.LeadBallsRecipient.ID,
			&m.DiscussionDate, &m.Book.ID, &m.Book.Title, &m.Book.Author, &m.Book.PublishedDate, &m.Book.PageCount,
			&m.Book.AverageRating, &m.Book.Thumbnail, &m.Book.Description, &m.Book.Link,
		); err != nil {
			return nil, err
		}
		meetings = append(meetings, m)
	}

	return meetings, nil
}

var (
	createQuery = `INSERT INTO meetings(bookID, jesus, jesusChooserID, 
										jesusExplanation, bookChooserID, goldenExplanation,
										goldenRecipientID, leadExplanation, leadRecipientID,
										discussionDate)
									 VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
									 RETURNING id`

	getAllQuery = `SELECT m.*, b.*
							FROM meetings AS m
							JOIN books AS b ON m.bookId = b.id`
)
