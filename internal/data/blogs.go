package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Blogs struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Year     int32  `json:"year"`
	Author   string `json:"author"`
	Content  string `json:"content"`

	CreatedAt time.Time `json:"-"`
	Version   string    `json:"version"`
}

type BlogsModel struct {
	DB *sql.DB
}

// Insert Add a placeholder method for inserting a new record in the movies table.
func (m BlogsModel) Insert(blog *Blogs) error {
	query := `
      INSERT INTO blogs(title, subtitle, year, author, content)
      VALUES ($1, $2, $3, $4, $5)
      RETURNING id, created_at, version`

	args := []interface{}{blog.Title, blog.Subtitle, blog.Year, blog.Author, blog.Content}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&blog.ID, &blog.CreatedAt, &blog.Version)
}

// Get Add a placeholder method for fetching a specific record from the movies table.
func (m BlogsModel) Get(id int64) (*Blogs, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
        SELECT id, created_at, title, year, content, subtitle,author, version
        FROM blogs
        WHERE id = $1`

	var blog Blogs

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&blog.ID,
		&blog.CreatedAt,
		&blog.Title,
		&blog.Year,
		&blog.Content,
		&blog.Subtitle,
		&blog.Author,
		&blog.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &blog, nil
}

// Update Add a placeholder method for updating a specific record in the movies table.
func (m BlogsModel) Update(blog *Blogs) error {
	return nil
}

// Delete Add a placeholder method for deleting a specific record from the movies table.
func (m BlogsModel) Delete(id int64) error {
	return nil
}
