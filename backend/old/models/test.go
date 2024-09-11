package models

import (
	"fmt"
	"log"
	"github.com/KUNSALISA/Login_SA/database"
    "github.com/KUNSALISA/Login_SA/models"
    "github.com/gin-gonic/gin"
)

func main() {
	// เชื่อมต่อกับฐานข้อมูล
	database.ConnectDatabase()

	// ค้นหาผู้ใช้
	var user models.User
	err := database.DB.First(&user, "email = ?", "john.doe@example.com").Error
	if err != nil {
		log.Fatalf("Error finding user: %v", err)
	}

	// ตรวจสอบรหัสผ่าน
	isValid := user.CheckPassword("yourpassword123")
	if isValid {
		fmt.Println("Password is valid")
	} else {
		fmt.Println("Password is invalid")
	}
}
