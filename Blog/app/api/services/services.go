package services

import (
	"blog-backend/app/db"
	"blog-backend/app/db/models"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthResponse struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Token string `json:"token"`
}

func SignUp(name, password string) (AuthResponse, error) {
	var user models.User

	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user.Name = name
	user.Password = string(bytes)

	alreadyExists := db.DB.Find(&user, name)

	if alreadyExists != nil {
		return AuthResponse{}, fmt.Errorf("username already exists")
	}

	err := db.DB.Create(&user).Error

	if err != nil {
		return AuthResponse{}, fmt.Errorf("failed to create user")
	}

	token, err := generateToken(user.ID) // Replace with your actual token generation logic
	if err != nil {
		log.Println("Error generating token:", err)
		return AuthResponse{}, err
	}

	return AuthResponse{
		Name:  user.Name,
		Type:  "jwt",
		Token: token,
	}, nil

}

func Login(name, password string) (AuthResponse, error) {
	var user models.User

	err := db.DB.Where("name = ?", name).First(&user).Error

	if err != nil {
		return AuthResponse{}, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return AuthResponse{}, fmt.Errorf("wrong password")
	}

	token, err := generateToken(user.ID)

	return AuthResponse{
		Name:  user.Name,
		Type:  "jwt",
		Token: token,
	}, nil

}

func CreateBlog(Id int, title string, description string, body string) (bool, error) {

	log.Println("===== CreateBlog DEBUG START =====")
	log.Printf("Input -> Id=%d, title=%q, description=%q, body=%q\n", Id, title, description, body)

	var user models.User
	err := db.DB.Where("id = ?", Id).First(&user).Error

	if err != nil {
		log.Printf("User lookup failed for Id=%d | Error=%v\n", Id, err)
		return false, fmt.Errorf("user not found")
	}

	log.Printf("User found -> ID=%d, Name=%q\n", user.ID, user.Name)

	blog := models.Blog{
		UserID:      user.ID,
		Title:       title,
		Description: description,
		Body:        body,
	}

	log.Printf("Blog to insert -> %+v\n", blog)

	err = db.DB.Create(&blog).Error

	if err != nil {
		log.Printf("Blog insert failed | Error=%v\n", err)
		return false, fmt.Errorf("server error")
	}

	return true, nil

}

func generateToken(userID uint) (string, error) {
	// Implement token generation logic (e.g., JWT)
	return "generated_token_here", nil
}
