package users

import "github.com/rtravitz/culture_knights/db"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u *User) Get(db *db.DB) error {
	return db.QueryRow("SELECT name FROM users WHERE id=$1",
		u.ID).Scan(&u.Name)
}

func (u *User) Update(db *db.DB) error {
	_, err := db.Exec("UPDATE users SET name=$1 WHERE id=$2",
		u.Name, u.ID)

	return err
}

func (u *User) Delete(db *db.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", u.ID)

	return err
}

func (u *User) Create(db *db.DB) error {
	err := db.QueryRow("INSERT INTO users(name) VALUES($1) RETURNING id",
		u.Name).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetUsers(db *db.DB, start, count int) ([]User, error) {
	rows, err := db.Query("SELECT id, name FROM users LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User

		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
