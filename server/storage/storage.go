package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dbFile string) error {
	var err error
	db, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}

	// Create the table if it doesn't exist
	createTable := `
        CREATE TABLE IF NOT EXISTS ShortURL (
            ShortURL TEXT UNIQUE,
            OriginalURL TEXT PRIMARY KEY
        );
    `
	_, err = db.Exec(createTable)
	if err != nil {
		return err
	}

	return nil
}

func SaveUrl(newUrl ShortURL) error {
	_, err := db.Exec("INSERT INTO ShortURL (ShortURL, OriginalURL) VALUES (?, ?)", newUrl.ShortURL, newUrl.OriginalURL)
	return err
}

func FindUrls() ([]ShortURL, error) {
	rows, err := db.Query("SELECT ShortURL, OriginalURL FROM ShortURL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []ShortURL
	for rows.Next() {
		var url ShortURL
		if err := rows.Scan(&url.ShortURL, &url.OriginalURL); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}

func FindLongUrl(shortUrl string) (string, error) {
	var originalUrl string
	err := db.QueryRow("SELECT OriginalURL FROM ShortURL WHERE ShortURL = ?", shortUrl).Scan(&originalUrl)
	if err != nil {
		return "", err
	}
	return originalUrl, nil
}
