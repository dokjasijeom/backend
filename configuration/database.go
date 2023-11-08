package configuration

import (
	"database/sql"
	"github.com/dokjasijeom/backend/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	dsn := os.Getenv("DEV_DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	// Migrate the schema
	database.AutoMigrate(&entity.User{})

	return database
}

func TestDataBase() {
	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to PlanetScale
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}
	defer rows.Close()

	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		log.Println(tableName)
	}

	// Connected success log
	log.Println("Connected to PlanetScale")

	defer db.Close()
}
