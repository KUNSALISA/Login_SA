package database

import (
	"fmt"
	"log"

	"github.com/KUNSALISA/Login_SA/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// เชื่อมต่อกับฐานข้อมูล
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}

	// สร้างผู้ใช้ใหม่
	user := models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	// แฮชรหัสผ่าน
	err = user.SetPassword("yourpassword123")
	if err != nil {
		log.Fatalf("Error setting password: %v", err)
	}

	// บันทึกผู้ใช้
	result := DB.Create(&user)
	if result.Error != nil {
		log.Fatalf("Error creating user: %v", result.Error)
	}

	fmt.Printf("User created: %+v\n", user)
}
