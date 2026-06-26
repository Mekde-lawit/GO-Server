package services

import (
	"encoding/json"
	"net/http"
	"udemygo/models"

	"github.com/jmoiron/sqlx"
)

var dbConn *sqlx.DB

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts := models.GetPosts()

	sqlStmt := `SELECT * FROM posts`
	rows, err := dbConn.Queryx(sqlStmt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tempPost models.Post
		if err := rows.StructScan(&tempPost); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, tempPost)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&posts)
}

func SetDB(db *sqlx.DB) {
	dbConn = db
}