package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
	"log"
)

func NewSeriesRepositoryImpl(DB *gorm.DB) repository.SeriesRepository {
	return &seriesRepositoryImpl{DB: DB}
}

type seriesRepositoryImpl struct {
	*gorm.DB
}

// Create Series
func (seriesRepository *seriesRepositoryImpl) CreateSeries(ctx context.Context, series entity.Series, model model.SeriesModel) (entity.Series, error) {
	var seriesResult entity.Series
	seriesResult = series
	var authorResult entity.SeriesAuthor

	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Create(&seriesResult)

	if model.GenreIds != nil {
		for _, genreId := range model.GenreIds {
			err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("Genres").Append(&entity.Genre{Id: genreId})
			if err != nil {
				log.Println("장르 연결 실패")
				exception.PanicLogging(err)
				//return entity.Series{}, err
			}
		}
	}

	if model.ProviderIds != nil {
		for _, providerId := range model.ProviderIds {
			err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("Providers").Append(&entity.Provider{Id: providerId})
			if err != nil {
				log.Println("제공자 연결 실패")
				exception.PanicLogging(err)
				//return entity.Series{}, err
			}
		}
	}

	if model.PublishDayIds != nil {
		for _, publishDayId := range model.PublishDayIds {
			err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("PublishDays").Append(&entity.PublishDay{Id: publishDayId})
			if err != nil {
				log.Println("연재 요일 연결 실패")
				exception.PanicLogging(err)
				//return entity.Series{}, err
			}
		}
	}

	if model.PublisherIds != nil {
		for _, publisherId := range model.PublisherIds {
			err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("Publishers").Append(&entity.Publisher{Id: publisherId})
			if err != nil {
				log.Println("출판사 연결 실패")
				exception.PanicLogging(err)
				//return entity.Series{}, err
			}
		}
	}

	if model.AuthorId != 0 {
		result := seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.AuthorId, PersonType: "author"})

		if result.Error != nil {
			log.Println("작가 연결 실패")
			exception.PanicLogging(result.Error)
		}
	}

	if model.IllustratorId != 0 {
		result := seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.IllustratorId, PersonType: "illustrator"})

		if result.Error != nil {
			log.Println("그림 작가 연결 실패")
			exception.PanicLogging(result.Error)
		}
	}

	if model.OriginalAuthorId != 0 {
		result := seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.OriginalAuthorId, PersonType: "original_author"})

		if result.Error != nil {
			log.Println("원작 작가 연결 실패")
			exception.PanicLogging(result.Error)
		}
	}

	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return entity.Series{}, result.Error
	}
	return seriesResult, nil
}

// Update series hash id
func (seriesRepository *seriesRepositoryImpl) UpdateSeriesHashId(ctx context.Context, id uint, hashId string) error {
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Where("id = ?", id).Updates(map[string]interface{}{"hash_id": hashId})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (seriesRepository *seriesRepositoryImpl) UpdateSeriesById(ctx context.Context, id uint, series entity.Series, model model.SeriesModel) (entity.Series, error) {
	var seriesResult entity.Series
	seriesResult = series
	var authorResult entity.SeriesAuthor

	if model.PublisherIds != nil {
		// already exist publishers all remove for gorm
		err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("Publishers").Clear()
		if err != nil {
			log.Println("출판사 연결 해제 실패")
			exception.PanicLogging(err)
		}

		seriesResult.Publishers = nil
		for _, publisherId := range model.PublisherIds {
			err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("Publishers").Append(&entity.Publisher{Id: publisherId})
			if err != nil {
				log.Println("출판사 연결 실패")
				exception.PanicLogging(err)
			}
		}
	}

	if model.GenreIds != nil {
		// already exist genres all remove for gorm
		err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("Genres").Clear()
		if err != nil {
			log.Println("장르 연결 해제 실패")
			exception.PanicLogging(err)
		}

		seriesResult.Genres = nil
		for _, genreId := range model.GenreIds {
			err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("Genres").Append(&entity.Genre{Id: genreId})
			if err != nil {
				log.Println("장르 연결 실패")
				exception.PanicLogging(err)
			}
		}
	}

	if model.PublishDayIds != nil {
		// already exist publish days all remove for gorm
		err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("PublishDays").Clear()
		if err != nil {
			log.Println("연재 요일 연결 해제 실패")
			exception.PanicLogging(err)
		}

		seriesResult.PublishDays = nil
		for _, publishDayId := range model.PublishDayIds {
			err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("PublishDays").Append(&entity.PublishDay{Id: publishDayId})
			if err != nil {
				log.Println("연재 요일 연결 실패")
				exception.PanicLogging(err)
			}
		}
	}

	if len(seriesResult.SeriesAuthors) > 0 {
		for _, author := range seriesResult.SeriesAuthors {
			if model.AuthorId != 0 && author.PersonId != model.AuthorId {
				result := seriesRepository.DB.WithContext(ctx).Model(&authorResult).Where("id = ?", author.Id).Delete(&entity.SeriesAuthor{})
				result = seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.AuthorId, PersonType: "author"})

				if result.Error != nil {
					log.Println("작가 삭제 실패")
				}
			}
			if model.IllustratorId != 0 && author.PersonId != model.IllustratorId {
				result := seriesRepository.DB.WithContext(ctx).Model(&authorResult).Where("id = ?", author.Id).Delete(&entity.SeriesAuthor{})
				result = seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.IllustratorId, PersonType: "illustrator"})

				if result.Error != nil {
					log.Println("그림 작가 삭제 실패")
				}
			}

			if model.OriginalAuthorId != 0 && author.PersonId != model.OriginalAuthorId {
				result := seriesRepository.DB.WithContext(ctx).Model(&authorResult).Where("id = ?", author.Id).Delete(&entity.SeriesAuthor{})

				if result.Error != nil {
					log.Println("원작 작가 삭제 실패")
				}
			}

			result := seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.AuthorId, PersonType: "author"})
			result = seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.IllustratorId, PersonType: "illustrator"})
			result = seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.OriginalAuthorId, PersonType: "original_author"})

			if result.Error != nil {
				log.Println("작가 업데이트 실패")
			}
		}
	} else {
		if model.AuthorId != 0 {
			result := seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.AuthorId, PersonType: "author"})
			if result.Error != nil {
				log.Println("작가 연결 실패")
				exception.PanicLogging(result.Error)
			}
		}

		if model.IllustratorId != 0 {
			result := seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.IllustratorId, PersonType: "illustrator"})
			if result.Error != nil {
				log.Println("그림 작가 연결 실패")
				exception.PanicLogging(result.Error)
			}
		}

		if model.OriginalAuthorId != 0 {
			result := seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.OriginalAuthorId, PersonType: "original_author"})
			if result.Error != nil {
				log.Println("원작 작가 연결 실패")
				exception.PanicLogging(result.Error)
			}
		}
	}

	if model.ProviderIds != nil {
		// already exist providers all remove for gorm
		err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("Providers").Clear()
		if err != nil {
			log.Println("제공자 연결 해제 실패")
			exception.PanicLogging(err)
		}

		seriesResult.Providers = nil
		for _, providerId := range model.ProviderIds {
			err := seriesRepository.DB.WithContext(ctx).Model(&seriesResult).Association("Providers").Append(&entity.Provider{Id: providerId})
			if err != nil {
				log.Println("제공자 연결 실패")
				exception.PanicLogging(err)
			}
		}
	}

	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Where("id = ?", id).Updates(&seriesResult)
	if result.Error != nil {
		exception.PanicLogging(result.Error)
		return entity.Series{}, result.Error
	}

	return seriesResult, nil
}

func (seriesRepository *seriesRepositoryImpl) DeleteSeriesById(ctx context.Context, id uint) error {
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Delete(&entity.Series{}, id)
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesById(ctx context.Context, id uint) (entity.Series, error) {
	var seriesResult entity.Series
	result := seriesRepository.DB.WithContext(ctx).Preload("Providers").Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("Episodes").First(&seriesResult, id)
	if result.RowsAffected == 0 {
		return entity.Series{}, nil
	}

	if seriesResult.SeriesType == "webnovel" {
		seriesResult.DisplayTags = "#웹소설 "
	} else {
		seriesResult.DisplayTags = "#웹툰 "
	}

	for genreI := range seriesResult.Genres {
		seriesResult.DisplayTags += "#" + seriesResult.Genres[genreI].Name + " "
	}

	seriesResult.TotalEpisode = uint(len(seriesResult.Episodes))

	// DisplayTags 마지막 공백 제거
	seriesResult.DisplayTags = seriesResult.DisplayTags[:len(seriesResult.DisplayTags)-1]

	// 작가 유형 반영해서 Authors 필드에 반영
	seriesResult.Authors = make([]entity.Person, 0)
	for _, sa := range seriesResult.SeriesAuthors {
		sa.Person.PersonType = sa.PersonType
		seriesResult.Authors = append(seriesResult.Authors, sa.Person)
	}

	return seriesResult, nil
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByPublishDayAndSeriesType(ctx context.Context, publishDay, seriesType string) ([]entity.Series, error) {
	var seriesResult []entity.Series
	var publishDayResult entity.PublishDay
	var seriesIds []uint

	seriesRepository.DB.WithContext(ctx).Model(&entity.PublishDay{}).Where("day = ?", publishDay).First(&publishDayResult)
	seriesRepository.DB.WithContext(ctx).Model(&entity.SeriesPublishDay{}).Where("publish_day_id = ?", publishDayResult.Id).Pluck("series_id", &seriesIds)
	seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Where("series_type = ?", seriesType).Where("id in (?)", seriesIds).Preload("Providers").Preload("PublishDays").Preload("Genres").Preload("Publishers").Preload("SeriesAuthors").Preload("Episodes").Find(&seriesResult)

	// series 결과 목록에서 Id 필드값을 제거
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
		// DisplayTags 마지막 공백 제거
		seriesResult[i].DisplayTags = seriesResult[i].DisplayTags[:len(seriesResult[i].DisplayTags)-1]
	}

	return seriesResult, nil
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByPublishDayId(publishDayId uint) ([]entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByGenreId(genreId uint) ([]entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByPersonId(personId uint) ([]entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByPublisherId(publisherId uint) ([]entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesBySeriesType(seriesType string) ([]entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByTitle(title string) (entity.Series, error) {
	panic("implement me")
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByHashId(ctx context.Context, hashId string) (entity.Series, error) {
	var seriesResult entity.Series
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Preload("Providers").Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("Episodes").Where("hash_id = ?", hashId).First(&seriesResult)
	if result.Error != nil {
		return entity.Series{}, nil
	}

	if seriesResult.SeriesType == "webnovel" {
		seriesResult.DisplayTags = "#웹소설 "
	} else {
		seriesResult.DisplayTags = "#웹툰 "
	}

	for genreI := range seriesResult.Genres {
		seriesResult.DisplayTags += "#" + seriesResult.Genres[genreI].Name + " "
	}

	seriesResult.TotalEpisode = uint(len(seriesResult.Episodes))

	// DisplayTags 마지막 공백 제거
	seriesResult.DisplayTags = seriesResult.DisplayTags[:len(seriesResult.DisplayTags)-1]

	// 작가 유형 반영해서 Authors 필드에 반영
	seriesResult.Authors = make([]entity.Person, 0)
	for _, sa := range seriesResult.SeriesAuthors {
		sa.Person.PersonType = sa.PersonType
		seriesResult.Authors = append(seriesResult.Authors, sa.Person)
	}

	return seriesResult, nil
}

func (seriesRepository *seriesRepositoryImpl) GetAllSeries(ctx context.Context) ([]entity.Series, error) {
	var seriesResult []entity.Series

	err := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Preload("Providers").Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("Episodes").Find(&seriesResult)
	if err.Error != nil {
		exception.PanicLogging(err.Error)
		return nil, err.Error
	}

	for i := range seriesResult {
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
	}

	return seriesResult, nil
}

// Like Series
func (seriesRepository *seriesRepositoryImpl) LikeSeries(ctx context.Context, userId uint, seriesId uint) error {
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.UserLikeSeries{}).Create(&entity.UserLikeSeries{UserId: userId, SeriesId: seriesId})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
