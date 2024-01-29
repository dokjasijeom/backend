package backoffice

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
	"image"
	"image/jpeg"
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

func (controller BackofficeSeriesController) Route(app fiber.Router) {
	series := app.Group("/series")
	series.Post("/", controller.CreateSeries)
	series.Get("/", controller.GetAllSeries)
	series.Get("/:id", controller.GetSeriesById)
	series.Delete("/:id", controller.DeleteSeriesById)
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
	request.Thumbnail = filename

	result, err := controller.SeriesService.CreateSeries(ctx.Context(), request)

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    result,
	})
}

// Get Series by Id
func (controller BackofficeSeriesController) GetSeriesById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	result, err := controller.SeriesService.GetSeriesById(ctx.Context(), uint(id))
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

// Get All Series
func (controller BackofficeSeriesController) GetAllSeries(ctx *fiber.Ctx) error {
	result, err := controller.SeriesService.GetAllSeries(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    result,
	})
}

// Delete Series By Id
func (controller BackofficeSeriesController) DeleteSeriesById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return err
	}

	result := controller.SeriesService.DeleteSeriesById(ctx.Context(), uint(id))
	if result != nil {
		return result
	}

	return ctx.Status(fiber.StatusNoContent).JSON(model.GeneralResponse{
		Code:    fiber.StatusNoContent,
		Message: "Success",
		Data:    nil,
	})
}

func imageProcessing(ctx context.Context, buffer []byte, quality int) (string, error) {
	filename := strings.Replace(uuid.New().String(), "-", "", -1) + ".webp"

	img, _, err := image.Decode(bytes.NewReader(buffer))
	if err != nil {
		exception.PanicLogging(err)
	}
	var jpegBuffer bytes.Buffer
	err = jpeg.Encode(&jpegBuffer, img, &jpeg.Options{Quality: quality})
	if err != nil {
		exception.PanicLogging(err)
	}

	cld, _ := cloudinary.New()
	_, err = cld.Upload.Upload(ctx, jpegBuffer, uploader.UploadParams{
		PublicID: filename,
		Folder:   "series",
	})

	return "series/" + filename, err
}
