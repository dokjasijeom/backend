package backoffice

import (
	"bytes"
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/image/webp"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func NewBackofficeSeriesController(seriesService *service.SeriesService, episodeService *service.EpisodeService, config configuration.Config) *BackofficeSeriesController {
	return &BackofficeSeriesController{SeriesService: *seriesService, EpisodeService: *episodeService, Config: config}
}

type BackofficeSeriesController struct {
	service.SeriesService
	service.EpisodeService
	configuration.Config
}

func (controller BackofficeSeriesController) Route(app fiber.Router) {
	series := app.Group("/series")
	series.Post("/", controller.CreateSeries)
	series.Get("/", controller.GetAllSeries)
	series.Get("/:id", controller.GetSeriesById)
	series.Patch("/:id", controller.UpdateSeriesById)
	series.Delete("/:id", controller.DeleteSeriesById)
	series.Post("/:id/episodes", controller.CreateEpisode)
}

// Create Series
func (controller BackofficeSeriesController) CreateSeries(ctx *fiber.Ctx) error {
	var request = model.SeriesModel{}
	form, err := ctx.MultipartForm()
	if err != nil {
		exception.PanicLogging(err)
	}
	if title := form.Value["title"]; len(title) > 0 {
		request.Title = title[0]
	}
	if seriesType := form.Value["seriesType"]; len(seriesType) > 0 {
		request.SeriesType = entity.SeriesType(seriesType[0])
	}
	if description := form.Value["description"]; len(description) > 0 {
		request.Description = description[0]
	}
	if isbn := form.Value["isbn"]; len(isbn) > 0 {
		request.ISBN = isbn[0]
	}
	if ecn := form.Value["ecn"]; len(ecn) > 0 {
		request.ECN = ecn[0]
	}
	if isComplete := form.Value["isComplete"]; len(isComplete) > 0 {
		request.IsComplete, err = strconv.ParseBool(isComplete[0])
		if err != nil {
			request.IsComplete = false
		}
	}
	if personId := form.Value["personId"]; len(personId) > 0 {
		u64, _ := strconv.ParseInt(personId[0], 10, 0)
		request.PersonId = uint(u64)
	}
	if publisherId := form.Value["publisherId"]; len(publisherId) > 0 {
		u64, _ := strconv.ParseInt(publisherId[0], 10, 0)
		request.PublisherId = uint(u64)
	}

	if genreIds := form.Value["genreIds"]; len(genreIds) > 0 {
		for _, v := range genreIds {
			u64, _ := strconv.ParseInt(v, 10, 0)
			request.GenreIds = append(request.GenreIds, uint(u64))
		}
	}
	if providerIds := form.Value["providerIds"]; len(providerIds) > 0 {
		for _, v := range providerIds {
			u64, _ := strconv.ParseInt(v, 10, 0)
			request.ProviderIds = append(request.ProviderIds, uint(u64))
		}
	}
	if publishDayIds := form.Value["publishDayIds"]; len(publishDayIds) > 0 {
		for _, v := range publishDayIds {
			u64, _ := strconv.ParseInt(v, 10, 0)
			request.PublishDayIds = append(request.PublishDayIds, uint(u64))
		}
	}

	fileheader, err := ctx.FormFile("image")
	if err != nil {
		log.Println(0)
		exception.PanicLogging(err)
	}

	file, err := fileheader.Open()
	if err != nil {
		exception.PanicLogging(err)
	}

	buffer, err := io.ReadAll(file)
	if err != nil {
		exception.PanicLogging(err)
	}

	ext := fileheader.Filename[strings.LastIndex(fileheader.Filename, "."):]

	filename, err := imageProcessing(ctx.Context(), ext, buffer, 50)
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

// Update Series by Id
func (controller BackofficeSeriesController) UpdateSeriesById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	seriesResult, err := controller.SeriesService.GetSeriesById(ctx.Context(), uint(id))
	seriesModel := model.SeriesModel{}

	form, err := ctx.MultipartForm()
	if err != nil {
		exception.PanicLogging(err)
	}
	if title := form.Value["title"]; len(title) > 0 && seriesResult.Title != title[0] {
		seriesResult.Title = title[0]
	}
	if seriesType := form.Value["seriesType"]; len(seriesType) > 0 && seriesResult.SeriesType != entity.SeriesType(seriesType[0]) {
		seriesResult.SeriesType = entity.SeriesType(seriesType[0])
	}
	if description := form.Value["description"]; len(description) > 0 && seriesResult.Description != description[0] {
		seriesResult.Description = description[0]
	}
	if isbn := form.Value["isbn"]; len(isbn) > 0 && seriesResult.ISBN != isbn[0] {
		seriesResult.ISBN = isbn[0]
	}
	if ecn := form.Value["ecn"]; len(ecn) > 0 && seriesResult.ECN != ecn[0] {
		seriesResult.ECN = ecn[0]
	}

	if isComplete := form.Value["isComplete"]; len(isComplete) > 0 {
		seriesResult.IsComplete, err = strconv.ParseBool(isComplete[0])
		if err != nil {
			seriesResult.IsComplete = false
		}
	}

	if publisherId := form.Value["publisherId"]; len(publisherId) > 0 {
		u64, _ := strconv.ParseInt(publisherId[0], 10, 0)
		seriesModel.PublisherId = uint(u64)
	}

	if personId := form.Value["personId"]; len(personId) > 0 {
		u64, _ := strconv.ParseInt(personId[0], 10, 0)
		//seriesResult.Authors = nil
		seriesModel.PersonId = uint(u64)
	}

	if genreIds := form.Value["genreIds"]; len(genreIds) > 0 {
		//seriesResult.Genres = nil
		for _, v := range genreIds {
			u64, _ := strconv.ParseInt(v, 10, 0)
			seriesModel.GenreIds = append(seriesModel.GenreIds, uint(u64))
		}
	}
	if providerIds := form.Value["providerIds"]; len(providerIds) > 0 {
		//seriesResult.Providers = nil
		for _, v := range providerIds {
			u64, _ := strconv.ParseInt(v, 10, 0)
			seriesModel.ProviderIds = append(seriesModel.ProviderIds, uint(u64))
		}
	}
	if publishDayIds := form.Value["publishDayIds"]; len(publishDayIds) > 0 {
		//seriesResult.PublishDays = nil
		for _, v := range publishDayIds {
			u64, _ := strconv.ParseInt(v, 10, 0)
			seriesModel.PublishDayIds = append(seriesModel.PublishDayIds, uint(u64))
		}
	}

	// if thumbnail is not empty, remove the old thumbnail
	if fileheader, err := ctx.FormFile("image"); err == nil {
		removeImage(ctx.Context(), seriesResult.Thumbnail)
		if err != nil {
			log.Println(0)
			exception.PanicLogging(err)
		}

		file, err := fileheader.Open()
		if err != nil {
			exception.PanicLogging(err)
		}

		buffer, err := io.ReadAll(file)
		if err != nil {
			exception.PanicLogging(err)
		}

		ext := fileheader.Filename[strings.LastIndex(fileheader.Filename, "."):]

		filename, err := imageProcessing(ctx.Context(), ext, buffer, 50)
		if err != nil {
			exception.PanicLogging(err)
		}
		seriesResult.Thumbnail = filename
	}

	result, err := controller.SeriesService.UpdateSeriesById(ctx.Context(), uint(id), seriesResult, seriesModel)

	return ctx.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
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

	for i, v := range result {
		result[i].Thumbnail = controller.Config.Get("CLOUDINARY_URL") + v.Thumbnail
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

func (controller BackofficeSeriesController) CreateEpisode(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)

	if err != nil {
		return err
	}

	var request struct {
		EpisodeNumber uint `json:"episodeNumber" validate:"required"`
	}
	err = ctx.BodyParser(&request)

	println(request.EpisodeNumber)

	//request.EpisodeNumber = strings.TrimSpace(request.EpisodeNumber)
	//// request.episodenumber to uint
	//episodeNumber, err := strconv.ParseUint(request.EpisodeNumber, 10, 0)
	//if err != nil {
	//	return err
	//}

	episodes, err := controller.EpisodeService.CreateBulkEpisode(ctx.Context(), uint(id), uint(request.EpisodeNumber))
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    fiber.StatusCreated,
		Message: "Success",
		Data:    episodes,
	})
}

func removeImage(ctx context.Context, filePath string) error {
	cld, err := configuration.NewCloudinaryConfigruation()
	if err != nil {
		exception.PanicLogging(err)
	}

	filename := strings.TrimPrefix(filePath, "series/")

	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: filename,
	})
	if err != nil {
		exception.PanicLogging(err)
	}

	return nil
}

func imageProcessing(ctx context.Context, fileExt string, buffer []byte, quality int) (string, error) {
	filename := strings.Replace(uuid.New().String(), "-", "", -1) + ".jpg"
	var img image.Image
	var err error

	if fileExt == ".webp" {
		img, err = webp.Decode(bytes.NewReader(buffer))
	} else {
		img, _, err = image.Decode(bytes.NewReader(buffer))
	}

	if err != nil {
		exception.PanicLogging(err)
	}
	var jpegBuffer bytes.Buffer
	err = jpeg.Encode(&jpegBuffer, img, &jpeg.Options{Quality: quality})
	if err != nil {
		exception.PanicLogging(err)
	}

	f, err := os.CreateTemp("./tmp/", "tempfile-")
	if err != nil {
		exception.PanicLogging(err)
	}

	if _, err := f.Write(jpegBuffer.Bytes()); err != nil {
		exception.PanicLogging(err)
	}

	cld, err := configuration.NewCloudinaryConfigruation()
	if err != nil {
		exception.PanicLogging(err)
	}

	filenameWithoutExt := strings.TrimSuffix(filename, ".jpg")

	uploadResult, err := cld.Upload.Upload(ctx, f.Name(), uploader.UploadParams{
		ResourceType: "image",
		PublicID:     filenameWithoutExt,
		Folder:       "series",
	})
	if err != nil {
		exception.PanicLogging(err)
	}

	defer f.Close()
	defer os.Remove(f.Name())

	log.Println(uploadResult)

	return "series/" + filename, err
}
