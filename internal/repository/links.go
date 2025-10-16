package repository

import (
	"database/sql"
	"time"
)

type Link struct {
	ID          int
	ShortURL    string
	OriginalURL string
	TimeCreated time.Time
}

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(database *sql.DB) *LinkRepository {
	return &LinkRepository{db: database}
}

func (r *LinkRepository) GetDB() *sql.DB {
	return r.db
}

func (r *LinkRepository) GetByOriginalURL(url string) (string, error) {
	var short_url string
	err := r.db.QueryRow("SELECT short_url FROM links WHERE original_url = $1", url).Scan(&short_url)
	return short_url, err
}

func (r *LinkRepository) GetByShortURL(short_url string) (string, error) {
	var url string
	err := r.db.QueryRow("SELECT original_url FROM links WHERE short_url = $1", short_url).Scan(&url)
	return url, err
}

func (r *LinkRepository) Insert(short_url, url string) error {
	_, err := r.db.Exec("INSERT INTO links (short_url, original_url) VALUES ($1, $2)", short_url, url)
	return err
}

func (r *LinkRepository) GetAll() ([]Link, error) {
	rows, err := r.db.Query("SELECT * FROM links")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []Link
	for rows.Next() {
		var l Link
		if err := rows.Scan(&l.ID, &l.ShortURL, &l.OriginalURL, &l.TimeCreated); err != nil {
			return nil, err
		}
		links = append(links, l)
	}
	return links, rows.Err()
}
