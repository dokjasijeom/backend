package configuration

import (
	"database/sql"
	"github.com/dokjasijeom/backend/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		//log.Fatal("failed to load env", err)
	}

	releaseMode := os.Getenv("RELEASE_MODE")
	log.Println("releaseMode: ", releaseMode)

	dsn := os.Getenv("DSN")

	tablePrefix := func(releaseMode string) string {
		if releaseMode == "development" {
			return "dev_"
		} else {
			return ""
		}
	}(releaseMode)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: tablePrefix,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	// Migrate the schema
	database.AutoMigrate(&entity.Role{})
	database.AutoMigrate(&entity.UserRole{})
	database.AutoMigrate(&entity.User{})
	database.AutoMigrate(&entity.Publisher{})
	database.AutoMigrate(&entity.PublishDay{})
	database.AutoMigrate(&entity.Genre{})
	database.AutoMigrate(&entity.Provider{})
	database.AutoMigrate(&entity.Person{})
	database.AutoMigrate(&entity.Series{})
	database.AutoMigrate(&entity.Episode{})
	database.AutoMigrate(&entity.SeriesPublisher{})
	database.AutoMigrate(&entity.SeriesAuthor{})
	database.AutoMigrate(&entity.SeriesGenre{})
	database.AutoMigrate(&entity.SeriesPublishDay{})
	database.AutoMigrate(&entity.SeriesProvider{})
	database.AutoMigrate(&entity.SeriesDailyView{})
	database.AutoMigrate(&entity.UserLikeSeries{})
	database.AutoMigrate(&entity.UserLikeSeriesCount{})

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
