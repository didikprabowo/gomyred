package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomyred/config"
	"github.com/gomyred/src/article/model"
	"log"
	"sync"
)

var (
	dbCon *sql.DB
)

// NewMySQL
func NewMySQL(cfg config.MySQL) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DSN())

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbCon = db

	return dbCon, nil

}

func GetDB() *sql.DB {
	return dbCon
}

// GetAllArticle
func GetAllArticle(jobs chan<- model.Article, wg *sync.WaitGroup, start int, end int) {
	defer wg.Done()
	q := fmt.Sprintf("SELECT id, title, body FROM news limit %d,%d",
		start, end)
	rows, err := dbCon.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var id int
	var title, body string

	for rows.Next() {
		if err := rows.Scan(&id, &title, &body); err != nil {
			log.Fatal(err)
		}
		jobs <- model.Article{ID: id, Title: title, Body: body}
	}
	close(jobs)
}
