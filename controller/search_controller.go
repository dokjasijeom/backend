package controller

import (
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/service"
	"github.com/gofiber/fiber/v2"
)

func NewSearchController(seriesService *service.SeriesService, config configuration.Config) *SearchController {
	return &SearchController{SeriesService: *seriesService, Config: config}
}

type SearchController struct {
	service.SeriesService
	configuration.Config
}

func (controller SearchController) Route(app fiber.Router) {
	search := app.Group("/search")
	search.Get("/autocomplete", controller.Autocomplete)
	search.Get("/", controller.SearchSeries)
}

func (controller SearchController) Autocomplete(ctx *fiber.Ctx) error {
	query := ctx.Query("query")
	if query == "" {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "success",
			"data":    []entity.Series{},
		})
	}

	result, err := controller.SeriesService.GetSeriesIdAndTitlesByTitle(ctx.Context(), query)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success",
		"data":    result,
	})
}

func (controller SearchController) SearchSeries(ctx *fiber.Ctx) error {
	query := ctx.Query("query")
	if query == "" {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "success",
			"data":    []entity.Series{},
		})
	}

	seriesResult, err := controller.SeriesService.GetSeriesByTitle(ctx.Context(), query)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
	}

	for i := range seriesResult {
		seriesResult[i].Id = 0

		if seriesResult[i].SeriesType == "webnovel" {
			seriesResult[i].DisplayTags = "#웹소설 "
		} else {
			seriesResult[i].DisplayTags = "#웹툰 "
		}

		for genreI := range seriesResult[i].Genres {
			seriesResult[i].DisplayTags += "#" + seriesResult[i].Genres[genreI].Name + " "
		}
		seriesResult[i].TotalEpisode = uint(len(seriesResult[i].Episodes))
		seriesResult[i].DisplayTags = seriesResult[i].DisplayTags[:len(seriesResult[i].DisplayTags)-1]

		// 작가 유형 반영해서 Authors 필드에 반영
		seriesResult[i].Authors = make([]entity.Person, 0)
		for _, sa := range seriesResult[i].SeriesAuthors {
			sa.Person.PersonType = sa.PersonType
			seriesResult[i].Authors = append(seriesResult[i].Authors, sa.Person)
		}
		// 제공자 정보를 Providers 필드에 반영
		seriesResult[i].Providers = make([]entity.Provider, 0)
		for _, sp := range seriesResult[i].SeriesProvider {
			sp.Provider.Link = sp.Link
			seriesResult[i].Providers = append(seriesResult[i].Providers, sp.Provider)
		}

		// publishDays remove id, Displayorder
		for j, _ := range seriesResult[i].PublishDays {
			seriesResult[i].PublishDays[j].Id = 0
			seriesResult[i].PublishDays[j].DisplayOrder = 0
		}
		// publishers remove field id, description, homepageurl, series
		for j, _ := range seriesResult[i].Publishers {
			seriesResult[i].Publishers[j].Id = 0
			seriesResult[i].Publishers[j].Description = ""
			seriesResult[i].Publishers[j].HomepageUrl = ""
			seriesResult[i].Publishers[j].Series = nil
		}

		seriesResult[i].Thumbnail = controller.Config.Get("CLOUDINARY_URL") + seriesResult[i].Thumbnail
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success",
		"data":    seriesResult,
	})
}
