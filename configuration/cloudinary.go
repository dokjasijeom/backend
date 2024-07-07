package configuration

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func NewCloudinaryConfigruation() (*cloudinary.Cloudinary, error) {
	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		//log.Fatal("failed to load env", err)
	}

	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)

	var logger = logrus.New()
	cld.Logger.Writer = logger.WithField("source", "cloudinary")
	cld.Logger.SetLevel(2)

	return cld, err
}
