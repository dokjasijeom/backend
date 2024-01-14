package configuration

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
	"os"
)

func NewCloudinaryConfigruation() *cloudinary.Cloudinary {
	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		//log.Fatal("failed to load env", err)
	}

	cld, _ := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)

	return cld
}
