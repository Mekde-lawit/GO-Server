package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"udemygo/models"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var dbConn *sqlx.DB

func GetAllPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	posts := models.GetPosts()

	sqlStmt := `SELECT * FROM posts`
	rows, err := dbConn.Queryx(sqlStmt)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tempPost models.Post
		if err := rows.StructScan(&tempPost); err != nil { //sqlx's method for taking the current row's columns
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, tempPost)
	}

	if err := rows.Err(); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(&posts)
}

func CreatePost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var post models.Post

	if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStmt := `INSERT INTO posts (title, body) VALUES ($1, $2) RETURNING id`

	if err := dbConn.QueryRow(sqlStmt, post.Title, post.Body).Scan(&post.ID); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("New record...")

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(&post)
}

func GetPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	idParam := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(res, "invalid id", http.StatusBadRequest)
		return
	}

	var post models.Post
	sqlStmt := `SELECT * FROM posts WHERE id = $1`

	err = dbConn.Get(&post, sqlStmt, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(res, "post not found", http.StatusNotFound)
			return
		}
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(res).Encode(&post)
}

func UpdatePost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	idParam := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(res, "invalid id", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	sqlStmt := `UPDATE posts SET title = $1, body = $2 WHERE id = $3`

	result, err := dbConn.Exec(sqlStmt, post.Title, post.Body, id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(res, "post not found", http.StatusNotFound)
		return
	}

	post.ID = id
	json.NewEncoder(res).Encode(&post)
}

func DeletePost(res http.ResponseWriter, req *http.Request) {
	idParam := mux.Vars(req)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(res, "invalid id", http.StatusBadRequest)
		return
	}

	sqlStmt := `DELETE FROM posts WHERE id = $1`

	result, err := dbConn.Exec(sqlStmt, id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(res, "post not found", http.StatusNotFound)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}

func SetDB(db *sqlx.DB) {
	dbConn = db
}
