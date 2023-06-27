package db

import (
	"fmt"
	"log"
	"os"

	"github.com/andy10134/asisPay-backend/models/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	DbHost := os.Getenv("HOST")
	DbName := os.Getenv("DBNAME")
	DbUsername := os.Getenv("USER")
	DbPassword := os.Getenv("PASSWORD")

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", DbHost, DbUsername, DbPassword, DbName)
	var db, err = gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Conexión fallida ;(")
	}

	DB = db
	fmt.Println("Conexión establecida :)")

	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Corriendo Migraciones")
	err = DB.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatal("Migración Fallida:  \n", err.Error())
		os.Exit(1)
	}
}
