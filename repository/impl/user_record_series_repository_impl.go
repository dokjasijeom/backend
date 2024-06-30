package impl

import (
	"context"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
	"log"
)

func NewUserRecordSeriesRepositoryImpl(DB *gorm.DB) repository.UserRecordSeriesRepository {
	return &userRecordSeriesRepositoryImpl{DB: DB}
}

type userRecordSeriesRepositoryImpl struct {
	*gorm.DB
}

// Get user record series by id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) GetUserRecordSeriesById(ctx context.Context, id uint) (entity.UserRecordSeries, error) {
	var userRecordSeries entity.UserRecordSeries
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("id = ?", id).Find(&userRecordSeries)
	return userRecordSeries, result.Error
}

// Get user record series by user id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) GetUserRecordSeriesByUserId(ctx context.Context, userId uint) ([]entity.UserRecordSeries, error) {
	var userRecordSeries []entity.UserRecordSeries
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("user_id = ?", userId).Find(&userRecordSeries)
	return userRecordSeries, result.Error
}

// Get user record series by user id and series id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) GetUserRecordSeriesByUserIdAndSeriesId(ctx context.Context, userId, seriesId uint) (entity.UserRecordSeries, error) {
	var userRecordSeries entity.UserRecordSeries
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("user_id = ? AND series_id = ?", userId, seriesId).First(&userRecordSeries)
	return userRecordSeries, result.Error
}

// Get user record series by user id and user record series id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) GetUserRecordSeriesByUserIdAndId(ctx context.Context, userId, id uint) (entity.UserRecordSeries, error) {
	var userRecordSeries entity.UserRecordSeries
	config := configuration.New()
	releaseMode := config.Get("RELEASE_MODE")

	tablePrefix := func(releaseMode string) string {
		if releaseMode == "development" {
			return "dev_"
		} else {
			return ""
		}
	}(releaseMode)
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("user_id = ? AND id = ?", userId, id).Preload("RecordEpisodes", func(db *gorm.DB) *gorm.DB {
		return db.Order(tablePrefix + "user_record_series_episodes.episode_number asc")
	}).First(&userRecordSeries)
	err := result.Error

	if userRecordSeries.SeriesId != 0 {
		series, err := userRecordSeriesRepository.GetUserRecordSeriesItem(ctx, userRecordSeries.SeriesId)
		if err != nil {
			return userRecordSeries, err
		}
		userRecordSeries.Series = series

		userRecordSeries.Series.Id = 0
		if userRecordSeries.Series.SeriesType == "webnovel" {
			userRecordSeries.Series.DisplayTags = "#웹소설 "
		} else {
			userRecordSeries.Series.DisplayTags = "#웹툰 "
		}

		for genreI := range userRecordSeries.Series.Genres {
			userRecordSeries.Series.Genres[genreI].Id = 0
			userRecordSeries.Series.DisplayTags += "#" + userRecordSeries.Series.Genres[genreI].Name + " "
		}
		userRecordSeries.Series.TotalEpisode = uint(len(userRecordSeries.Series.Episodes))
		userRecordSeries.Series.DisplayTags = userRecordSeries.Series.DisplayTags[:len(userRecordSeries.Series.DisplayTags)-1]

		userRecordSeries.Series.Authors = make([]entity.Person, 0)
		for _, sa := range userRecordSeries.Series.SeriesAuthors {
			sa.Person.Id = 0
			sa.Person.PersonType = sa.PersonType
			userRecordSeries.Series.Authors = append(userRecordSeries.Series.Authors, sa.Person)
		}

		userRecordSeries.Series.Providers = make([]entity.Provider, 0)
		for _, sp := range userRecordSeries.Series.SeriesProvider {
			sp.Provider.Id = 0
			sp.Provider.Link = sp.Link
			userRecordSeries.Series.Providers = append(userRecordSeries.Series.Providers, sp.Provider)
		}

		for j, _ := range userRecordSeries.Series.PublishDays {
			userRecordSeries.Series.PublishDays[j].Id = 0
			userRecordSeries.Series.PublishDays[j].DisplayOrder = 0
		}

		for j, _ := range userRecordSeries.Series.Publishers {
			userRecordSeries.Series.Publishers[j].Id = 0
			userRecordSeries.Series.Publishers[j].Description = ""
			userRecordSeries.Series.Publishers[j].HomepageUrl = ""
			userRecordSeries.Series.Publishers[j].Series = nil
		}

		userRecordSeries.Series.Episodes = nil
		userRecordSeries.Series.Thumbnail = config.Get("CLOUDINARY_URL") + userRecordSeries.Series.Thumbnail

	}

	return userRecordSeries, err
}

// Get user record series item
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) GetUserRecordSeriesItem(ctx context.Context, id uint) (*entity.Series, error) {
	var series *entity.Series
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("id = ?", id).Preload("Genres").Preload("SeriesAuthors.Person").Preload("SeriesProvider.Provider").Preload("PublishDays").Preload("Publishers").Preload("Episodes").Find(&series)
	return series, result.Error
}

// Create user record series
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) CreateUserRecordSeries(ctx context.Context, userRecordSeries entity.UserRecordSeries) (entity.UserRecordSeries, error) {
	log.Println("userRecordSeries: ", userRecordSeries)
	result := userRecordSeriesRepository.DB.WithContext(ctx).Create(&userRecordSeries)
	if userRecordSeries.SeriesId == 0 {
		userRecordSeries.Series = nil
	}
	userRecordSeries.RecordEpisodes = nil
	return userRecordSeries, result.Error
}

// Update user record series by id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) UpdateUserRecordSeriesByUserIdAndId(ctx context.Context, userId, id uint, userRecordSeries entity.UserRecordSeries) (entity.UserRecordSeries, error) {
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("id = ? and user_id = ?", id, userId).Updates(&userRecordSeries)
	return userRecordSeries, result.Error
}

// Delete user record series by id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) DeleteUserRecordSeriesByUserIdAndId(ctx context.Context, userId, id uint) error {
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("id = ? and user_id = ?", id, userId).Delete(&entity.UserRecordSeries{})
	return result.Error
}

// Delete user record series by user id and series id
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) DeleteUserRecordSeriesByUserIdAndSeriesId(ctx context.Context, userId, seriesId uint) error {
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("user_id = ? AND series_id = ?", userId, seriesId).Delete(&entity.UserRecordSeries{})
	return result.Error
}

// Update User Record Series
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) UpdateUserRecordSeries(ctx context.Context, userId, id uint, request entity.UserRecordSeries) (entity.UserRecordSeries, error) {
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("id = ? and user_id = ?", id, userId).Updates(&request)
	return request, result.Error
}

// Get User Complete Records
func (userRecordSeriesRepository *userRecordSeriesRepositoryImpl) GetUserCompleteRecords(ctx context.Context, userId uint) ([]entity.UserRecordSeries, error) {
	var userRecordSeries []entity.UserRecordSeries
	result := userRecordSeriesRepository.DB.WithContext(ctx).Where("user_id = ?", userId).Where("read_completed = ?", true).Preload("RecordEpisodes").Preload("Genres").Preload("SeriesAuthors.Person").Preload("SeriesProvider.Provider").Preload("PublishDays").Preload("Publishers").Find(&userRecordSeries)
	return userRecordSeries, result.Error
}
