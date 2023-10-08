package main

import (
	"fmt"
	"log"
	"os"
	"server/config"
	"server/models"
	"server/routers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var err error

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file")
	}
}
func main() {
	dsn := os.Getenv("DATABASE_URL")
	config.DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("statuse", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})

	route := routers.SetupRouter()
	route.Run()
}
