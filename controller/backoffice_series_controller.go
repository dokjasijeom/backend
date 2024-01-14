package controller

import (
	"bytes"
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/h2non/bimg"
	"io"
	"strings"
)

func NewBackofficeSeriesController(seriesService *service.SeriesService, config configuration.Config) *BackofficeSeriesController {
	return &BackofficeSeriesController{SeriesService: *seriesService, Config: config}
}

type BackofficeSeriesController struct {
	service.SeriesService
	configuration.Config
}

func (controller BackofficeSeriesController) Route(app *fiber.App) {
	app.Post("/series", controller.CreateSeries)
	app.Get("/series", controller.GetAllSeries)
	app.Get("/series/:id", controller.GetSeriesById)
	app.Delete("/series/:id", controller.DeleteSeriesById)
}

// Create Series
func (controller BackofficeSeriesController) CreateSeries(ctx *fiber.Ctx) error {
	var request = model.SeriesModel{}
	err := ctx.BodyParser(&request)
	fileheader, _ := ctx.FormFile("image")

	file, err := fileheader.Open()
	if err != nil {
		exception.PanicLogging(err)
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		exception.PanicLogging(err)
	}

	filename, err := imageProcessing(ctx.Context(), buffer, 50)
	if err != nil {
		exception.PanicLogging(err)
	}
}

func imageProcessing(ctx context.Context, buffer []byte, quality int) (string, error) {
	filename := strings.Replace(uuid.New().String(), "-", "", -1) + ".webp"
	converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
	if err != nil {
		return filename, err
	}

	processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: quality})
	if err != nil {
		return filename, err
	}

	cld, _ := cloudinary.New()
	_, err = cld.Upload.Upload(ctx, bytes.NewReader(processed), uploader.UploadParams{
		PublicID: filename,
		Folder:   "series",
	})

	return "series/" + filename, err
}
