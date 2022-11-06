package databases

import (
	"BE_Rakamin/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DBInit() *gorm.DB {
	err := godotenv.Load(".env")
	serverUrl := os.Getenv("SQL_SERVER")
	db, err := gorm.Open("mysql", serverUrl)
	if err != nil {
		log.Fatal("failed to connect to database = ", err.Error())
	}
	db.AutoMigrate(models.Person{})
	return db
}
