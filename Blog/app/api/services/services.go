package services

import "Blog/app/db/models"

type LoginUserResponse struct {
	models.User
	Type  string `json:"type"`
	Token string `json:"token"`
}
