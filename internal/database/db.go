package database

import (
	"database/sql"
	"log"
	"newsbot/internal/models"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func (db *Database) Close() {
	if db.Conn != nil {
		db.Conn.Close()
	}
}



func NewDatabase(host string, port string, user string, pass string, dbname string) (*Database, error) {
	connStr := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Database connection successful!")
	return &Database{Conn: db}, nil
}

func SelectRandomArticle(db *sql.DB) models.Article {
	log.Println("SelectRandomArticle start")
	var article models.Article

	query := "SELECT id, date, title, content, url FROM articles_habr ORDER BY RANDOM() LIMIT 1"
	row := db.QueryRow(query)
	err := row.Scan(&article.ID, &article.Date, &article.Title, &article.Content, &article.Url)
	if err != nil {
		log.Println(err)
		return models.Article{}
	}

	return article
}

func SelectTodayPosts(db *sql.DB) models.Article {
	var article models.Article
	currentTime := time.Now()
	roundedTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	formattedTime := roundedTime.Format("2006-01-02 15:04:05-07")
	log.Println(formattedTime)

	query := "SELECT * FROM articles_habr WHERE DATE_TRUNC('day', date) = $1"

	rows, err := db.Query(query, formattedTime)
	if err != nil {
		log.Println("Error query ", err)
		return article
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&article.ID, &article.Date, &article.Title, &article.Content, &article.Url)
		if err != nil {
			log.Println(err)
			return models.Article{}
		}
	}
	return article
}
