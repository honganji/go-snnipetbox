package models

import (
	"database/sql"
	"errors"
	"time"
)

// Snippet represents a snippet record in the database
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// / Insert the snippet into the database and return the ID of the new snippet.
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// / Retrieve the snippet with the given ID from the database.
func (m *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := m.DB.QueryRow(stmt, id)
	var s Snippet
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}
	return s, nil
}

// / Retrieve the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var snippets []Snippet
	for rows.Next() {
		var s Snippet
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
