package main

import (
	"fmt"
	"log"
    "github.com/KUNSALISA/Login_SA/database"
    "github.com/KUNSALISA/Login_SA/handlers"
    "github.com/gin-gonic/gin"

)

func main() {
    r := gin.Default()

    // Connect to database
    database.ConnectDatabase()

    // Routes
    r.POST("/login", handlers.Login)
    r.POST("/register", handlers.Register)

    // Start the server
    r.Run(":8080") // Run on port 8080

	fmt.Println("Database connected, schema migrated, and initial data created.")

	var user models.User
	err := database.DB.First(&user, "email = ?", "john.doe@example.com").Error
	if err != nil {
		log.Fatalf("Error finding user: %v", err)
	}

	isValid := user.CheckPassword("yourpassword123")
	if isValid {
		fmt.Println("Password is valid")
	} else {
		fmt.Println("Password is invalid")
	}
}




