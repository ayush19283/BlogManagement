package main

import (
	"blog-backend/app/api/router"
	"blog-backend/app/db"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := db.Init(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	if err := db.Migrate(); err != nil {
		log.Fatalf("Error creating Tables: %v", err)
	}

	fmt.Println("Application started successfully!")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.SetupRouters(r)
	r.Run()

}
