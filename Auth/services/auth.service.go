package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gorilla/mux"
)

func Register (w http.ResponseWriter, r *http.Request){
	w.Header().Set()
	var user models.User
}



func Login (w http.ResponseWriter, r *http.Request){}

func Logout (w http.ResponseWriter, r *http.Request){}

func Protected (w http.ResponseWriter, r *http.Request){}