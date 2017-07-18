package users

import (
	"database/sql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u *User) getUser(db *sql.DB) error {
	return db.QueryRow("SELECT name FROM users WHERE id=$1",
		u.ID).Scan(&u.Name)
}

func (u *User) updateUser(db *sql.DB) error {
	_, err := db.Exec("UPDATE users SET name=$1 WHERE id=$2",
		u.Name, u.ID)

	return err
}

func (u *User) deleteUser(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", u.ID)

	return err
}

func (u *User) createUser(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO users(name) VALUES($1) RETURNING id",
		u.Name).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func getUsers(db *sql.DB, start, count int) ([]User, error) {
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
