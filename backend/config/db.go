package config

import (
	"fmt"

	"github.com/KUNSALISA/Login_SA/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {

	db.AutoMigrate(
		&entity.Member{},
	)

	hashedPassword, _ := HashPassword("123456")

	User := &entity.Member{
		Username: "Software",
		Email:     "sa@gmail.com",
		Password:  hashedPassword,
		
	}
	db.FirstOrCreate(User, &entity.Member{
		Email: "sa@gmail.com",
	})

}