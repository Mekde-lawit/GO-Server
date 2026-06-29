package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"auth/models"
	"auth/utils"

	"gorm.io/gorm"
)
var secretKey = os.Getenv("JWT_SECRET")
var DB *gorm.DB

func Register (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
     if user.Email == "" || user.Name == "" || user.Password == "" {
    http.Error(w, "All Fields Are Required", http.StatusBadRequest)
    return
        }

	var existingUser models.User
    result := DB.Where("email = ?", user.Email).First(&existingUser)

	 if result.Error == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	 }else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Database error: "+result.Error.Error(), http.StatusInternalServerError)
		return
	 }

	hashedPassword := utils.HashPassword(user.Password)
	if hashedPassword == nil {
        http.Error(w, "Could not hash password", http.StatusInternalServerError)
        return
    }

	user.Password = string(hashedPassword)
  // Save user
    result = DB.Create(&user)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }	
       
    tokenString, err := utils.GenerateJWT(user, secretKey)
    if err != nil {
        http.Error(w, "Failed to generate token"+ err.Error(), http.StatusInternalServerError)
        return
    }

	w.WriteHeader(http.StatusCreated)
	response := models.UserResponse{
    ID:    user.ID,
    Name:  user.Name,
    Email: user.Email,
    Role:  user.Role,
    Token: tokenString,
}
json.NewEncoder(w).Encode(response)
}

func Login (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body: "+ err.Error(), http.StatusBadRequest)
        return
    }

    if user.Email == "" || user.Password == "" {
    http.Error(w, "All Fields Are Required", http.StatusBadRequest)
    return
        }

	DB = utils.GetDB()

		var existingUser models.User
        result := DB.Where("email = ?", user.Email).First(&existingUser)

	 if result.Error != nil || utils.CheckPassword(user.Password, existingUser.Password) != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }     
    tokenString, err := utils.GenerateJWT(existingUser, secretKey)
    if err != nil {
        http.Error(w, "Failed to generate token"+ err.Error(), http.StatusInternalServerError)
        return
    }
	w.WriteHeader(http.StatusOK)
	response := models.UserResponse{
    ID:    existingUser.ID,
    Name:  existingUser.Name,
    Email: existingUser.Email,
    Role:  existingUser.Role,
	Token: tokenString,
}
json.NewEncoder(w).Encode(response)
}

func Logout (w http.ResponseWriter, r *http.Request){}

func Protected (w http.ResponseWriter, r *http.Request){}