package services

import (
	"blog-backend/app/db"
	"blog-backend/app/db/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type SignUpUserResponse struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Token string `json:"token"`
}

func SignUp(name, password string) (SignUpUserResponse, error) {
	var user models.User

	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user.Name = name
	user.Password = string(bytes)

	err := db.DB.Create(&user).Error

	if err != nil {
		return SignUpUserResponse{}, nil
	}

	token, err := generateToken(user.ID) // Replace with your actual token generation logic
	if err != nil {
		log.Println("Error generating token:", err)
		return SignUpUserResponse{}, err
	}

	return SignUpUserResponse{
		Name:  user.Name,
		Type:  "jwt",
		Token: token,
	}, nil

}

func generateToken(userID uint) (string, error) {
	// Implement token generation logic (e.g., JWT)
	return "generated_token_here", nil
}
