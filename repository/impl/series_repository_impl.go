package impl

import (
	"context"
	"github.com/dokjasijeom/backend/configuration"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/exception"
	"github.com/dokjasijeom/backend/model"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
	"log"
	"time"
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
	var providerResult entity.SeriesProvider

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

	if model.Providers != nil {
		for _, provider := range model.Providers {
			result := seriesRepository.DB.WithContext(ctx).Model(&providerResult).Create(&entity.SeriesProvider{SeriesId: seriesResult.Id, ProviderId: provider.ProviderId, Link: provider.Link})
			if result.Error != nil {
				log.Println("제공자 연결 실패")
				exception.PanicLogging(result.Error)
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
	var providerResult entity.SeriesProvider

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
				seriesRepository.DB.WithContext(ctx).Model(&authorResult).Where("id = ?", author.Id).Delete(&entity.SeriesAuthor{})
				seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.AuthorId, PersonType: "author"})
			}
			if model.IllustratorId != 0 && author.PersonId != model.IllustratorId {
				seriesRepository.DB.WithContext(ctx).Model(&authorResult).Where("id = ?", author.Id).Delete(&entity.SeriesAuthor{})
				seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.IllustratorId, PersonType: "illustrator"})
			}

			if model.OriginalAuthorId != 0 && author.PersonId != model.OriginalAuthorId {
				seriesRepository.DB.WithContext(ctx).Model(&authorResult).Where("id = ?", author.Id).Delete(&entity.SeriesAuthor{})
				seriesRepository.DB.WithContext(ctx).Model(&authorResult).Create(&entity.SeriesAuthor{SeriesId: seriesResult.Id, PersonId: model.OriginalAuthorId, PersonType: "original_author"})
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

	if len(seriesResult.SeriesProvider) > 0 {
		if model.Providers != nil {
			for _, provider := range seriesResult.SeriesProvider {
				// find seriesResult provider array for providerid
				providerExist := false
				for _, p := range model.Providers {
					if provider.ProviderId == p.ProviderId {
						providerExist = true
						break
					}
				}

				if !providerExist {
					result := seriesRepository.DB.WithContext(ctx).Model(&providerResult).Create(&entity.SeriesProvider{SeriesId: seriesResult.Id, ProviderId: provider.ProviderId, Link: provider.Link})
					if result.Error != nil {
						log.Println("제공자 연결 실패")
						exception.PanicLogging(result.Error)
					}
				}
			}
		}
	} else {
		if model.Providers != nil {
			for _, provider := range model.Providers {
				result := seriesRepository.DB.WithContext(ctx).Model(&providerResult).Create(&entity.SeriesProvider{SeriesId: seriesResult.Id, ProviderId: provider.ProviderId, Link: provider.Link})
				if result.Error != nil {
					log.Println("제공자 연결 실패")
					exception.PanicLogging(result.Error)
				}
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
	result := seriesRepository.DB.WithContext(ctx).Preload("SeriesProvider.Provider").Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("Episodes").First(&seriesResult, id)
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

	// DisplayTags 마지막 공백 제거
	seriesResult.DisplayTags = seriesResult.DisplayTags[:len(seriesResult.DisplayTags)-1]

	// 작가 유형 반영해서 Authors 필드에 반영
	seriesResult.Authors = make([]entity.Person, 0)
	for _, sa := range seriesResult.SeriesAuthors {
		sa.Person.PersonType = sa.PersonType
		seriesResult.Authors = append(seriesResult.Authors, sa.Person)
	}
	// 제공자 정보를 Providers 필드에 반영
	seriesResult.Providers = make([]entity.Provider, 0)
	for _, sp := range seriesResult.SeriesProvider {
		sp.Provider.Link = sp.Link
		seriesResult.Providers = append(seriesResult.Providers, sp.Provider)
	}

	return seriesResult, nil
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByPublishDayAndSeriesType(ctx context.Context, publishDay, seriesType string, page, pageSize int) (model.SeriesWithPagination, error) {
	var seriesResult []entity.Series
	var publishDayResult entity.PublishDay
	var seriesIds []uint

	seriesRepository.DB.WithContext(ctx).Model(&entity.PublishDay{}).Where("day = ?", publishDay).First(&publishDayResult)
	seriesRepository.DB.WithContext(ctx).Model(&entity.SeriesPublishDay{}).Where("publish_day_id = ?", publishDayResult.Id).Pluck("series_id", &seriesIds)
	seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Scopes(configuration.Paginate(page, pageSize)).Where("series_type = ?", seriesType).Where("id in (?)", seriesIds).Preload("SeriesProvider.Provider").Preload("PublishDays").Preload("Genres").Preload("Publishers").Preload("SeriesAuthors.Person").Preload("Episodes").Find(&seriesResult)

	var totalCount int64
	seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Where("series_type = ?", seriesType).Where("id in (?)", seriesIds).Count(&totalCount)

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
	}

	// paignation 정보 추가
	// totalCount, currentPage, nextPage, pageSize, totalPage entity
	// totalCount: 전체 데이터 수
	// currentPage: 현재 페이지
	// hasNext: 다음 페이지가 있는지 여부
	// pageSize: 페이지당 데이터 수
	// totalPage: 전체 페이지 수
	totalPage := (int(totalCount) / pageSize) + 1
	hasNext := func() bool {
		if page >= totalPage {
			return false
		}
		return true
	}()

	SeriesWithPagination := model.SeriesWithPagination{
		Series: seriesResult,
		Pagination: model.Pagination{
			TotalCount:  int(totalCount),
			CurrentPage: page,
			HasNext:     hasNext,
			PageSize:    pageSize,
			TotalPage:   totalPage,
		},
	}

	return SeriesWithPagination, nil
}

func (seriesRepository *seriesRepositoryImpl) GetNewEpisodeUpdateProviderSeries(ctx context.Context, provider, seriesType string, page, pageSize int) (model.SeriesWithPagination, error) {
	var seriesResult []entity.Series
	var providerResult entity.Provider
	var seriesIds []uint
	var episodeIds []uint

	// 오늘 날짜 가져오기
	now := time.Now()
	// 오늘 날짜에서 하루를 빼서 어제 날짜 가져오기
	// 임시로 -30일로 설정
	yesterday := now.AddDate(0, 0, -30)
	// 어제 날짜의 시간을 0시 0분 0초로 맞추기
	yesterday = time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, time.Local)

	// episode 테이블에서 어제 날짜와 오늘 날짜의 데이터를 비교해서 새로운 에피소드가 있는 series id를 중복값 제거해서 가져온다.
	seriesRepository.DB.WithContext(ctx).Model(&entity.Episode{}).Where("created_at BETWEEN ? AND ?", yesterday.Format(time.DateTime), now.Format(time.DateTime)).Pluck("id", &episodeIds)
	seriesRepository.DB.WithContext(ctx).Model(&entity.SeriesEpisode{}).Where("episode_id in (?)", episodeIds).Distinct().Pluck("series_id", &seriesIds)
	seriesRepository.DB.WithContext(ctx).Model(&entity.Provider{}).Where("name = ?", provider).First(&providerResult)
	seriesRepository.DB.WithContext(ctx).Model(&entity.SeriesProvider{}).Where("series_id in (?) and provider_id = ?", seriesIds, providerResult.Id).Distinct().Pluck("series_id", &seriesIds)
	seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Where("id in (?) and series_type = ?", seriesIds, seriesType).Preload("SeriesProvider.Provider").Preload("PublishDays").Preload("Genres").Preload("Publishers").Preload("SeriesAuthors.Person").Find(&seriesResult)

	var totalCount int64
	seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Where("series_type = ?", seriesType).Where("id in (?)", seriesIds).Count(&totalCount)

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
	}

	// paignation 정보 추가
	// totalCount, currentPage, nextPage, pageSize, totalPage entity
	// totalCount: 전체 데이터 수
	// currentPage: 현재 페이지
	// hasNext: 다음 페이지가 있는지 여부
	// pageSize: 페이지당 데이터 수
	// totalPage: 전체 페이지 수
	totalPage := (int(totalCount) / pageSize) + 1
	hasNext := func() bool {
		if page >= totalPage {
			return false
		}
		return true
	}()

	SeriesWithPagination := model.SeriesWithPagination{
		Series: seriesResult,
		Pagination: model.Pagination{
			TotalCount:  int(totalCount),
			CurrentPage: page,
			HasNext:     hasNext,
			PageSize:    pageSize,
			TotalPage:   totalPage,
		},
	}

	return SeriesWithPagination, nil
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

func (seriesRepository *seriesRepositoryImpl) GetSeriesByTitle(ctx context.Context, title string) ([]entity.Series, error) {
	var seriesResult []entity.Series
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Where("title LIKE ?", "%"+title+"%").Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("Episodes").Find(&seriesResult)
	if result.Error != nil {
		return nil, result.Error
	}

	return seriesResult, nil
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesIdAndTitlesByTitle(ctx context.Context, title string) ([]entity.Series, error) {
	var seriesResults []entity.Series
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Where("title LIKE ?", "%"+title+"%").Find(&seriesResults)
	if result.Error != nil {
		return nil, result.Error
	}
	return seriesResults, nil
}

func (seriesRepository *seriesRepositoryImpl) GetSeriesByHashId(ctx context.Context, hashId string) (entity.Series, error) {
	var seriesResult entity.Series
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Preload("SeriesProvider.Provider").Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("Episodes").Where("hash_id = ?", hashId).First(&seriesResult)
	if result.Error != nil {
		return entity.Series{}, result.Error
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
	// 제공자 정보를 Providers 필드에 반영
	seriesResult.Providers = make([]entity.Provider, 0)
	for _, sp := range seriesResult.SeriesProvider {
		sp.Provider.Link = sp.Link
		seriesResult.Providers = append(seriesResult.Providers, sp.Provider)
	}

	return seriesResult, nil
}

func (seriesRepository *seriesRepositoryImpl) GetBackofficeAllSeries(ctx context.Context) ([]entity.Series, error) {
	var seriesResult []entity.Series

	err := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("Episodes").Find(&seriesResult)
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
		// 제공자 정보를 Providers 필드에 반영
		seriesResult[i].Providers = make([]entity.Provider, 0)
		for _, sp := range seriesResult[i].SeriesProvider {
			sp.Provider.Link = sp.Link
			seriesResult[i].Providers = append(seriesResult[i].Providers, sp.Provider)
		}
	}

	return seriesResult, nil
}

func (seriesRepository *seriesRepositoryImpl) GetAllSeries(ctx context.Context, page, pageSize int) (model.SeriesWithPagination, error) {
	var seriesResult []entity.Series

	err := seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Scopes(configuration.Paginate(page, pageSize)).Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("Episodes").Find(&seriesResult)
	if err.Error != nil {
		exception.PanicLogging(err.Error)
		return model.SeriesWithPagination{}, err.Error
	}
	var totalCount int64
	err = seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Count(&totalCount)
	if err.Error != nil {
		exception.PanicLogging(err.Error)
		return model.SeriesWithPagination{}, err.Error
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
		// 제공자 정보를 Providers 필드에 반영
		seriesResult[i].Providers = make([]entity.Provider, 0)
		for _, sp := range seriesResult[i].SeriesProvider {
			sp.Provider.Link = sp.Link
			seriesResult[i].Providers = append(seriesResult[i].Providers, sp.Provider)
		}
	}

	// paignation 정보 추가
	// totalCount, currentPage, nextPage, pageSize, totalPage entity
	// totalCount: 전체 데이터 수
	// currentPage: 현재 페이지
	// hasNext: 다음 페이지가 있는지 여부
	// pageSize: 페이지당 데이터 수
	// totalPage: 전체 페이지 수
	totalPage := (int(totalCount) / pageSize) + 1
	hasNext := func() bool {
		if page >= totalPage {
			return false
		}
		return true
	}()

	SeriesWithPagination := model.SeriesWithPagination{
		Series: seriesResult,
		Pagination: model.Pagination{
			TotalCount:  int(totalCount),
			CurrentPage: page,
			HasNext:     hasNext,
			PageSize:    pageSize,
			TotalPage:   totalPage,
		},
	}

	return SeriesWithPagination, nil
}

// Get All Category Series
func (seriesRepository *seriesRepositoryImpl) GetAllCategorySeries(ctx context.Context, seriesType entity.SeriesType, genre string, providers []string, page, pageSize int) (model.SeriesWithPagination, error) {
	var seriesResult []entity.Series
	var genreResult entity.Genre
	var seriesIds []uint

	if genre != "" {
		if len(providers) == 0 {
			seriesRepository.DB.WithContext(ctx).Model(&entity.Genre{}).Where("hash_id = ?", genre).First(&genreResult)
			seriesRepository.DB.WithContext(ctx).Model(&entity.SeriesGenre{}).Where("genre_id = ?", genreResult.Id).Pluck("series_id", &seriesIds)
			seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Scopes(configuration.Paginate(page, pageSize)).Where("id in (?)", seriesIds).Where("series_type = ?", seriesType).Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("SeriesProvider.Provider").Preload("Episodes").Find(&seriesResult)
		} else {
			//var providerResult entity.Provider
			var providerIds []uint
			seriesRepository.DB.WithContext(ctx).Model(&entity.Provider{}).Where("hash_id in (?)", providers).Pluck("id", &providerIds)
			seriesRepository.DB.WithContext(ctx).Model(&entity.Genre{}).Where("hash_id = ?", genre).First(&genreResult)
			seriesRepository.DB.WithContext(ctx).Model(&entity.SeriesGenre{}).Where("genre_id = ?", genreResult.Id).Pluck("series_id", &seriesIds)
			seriesRepository.DB.WithContext(ctx).Model(&entity.SeriesProvider{}).Where("provider_id in (?)", providerIds).Where("series_id in (?)", seriesIds).Pluck("series_id", &seriesIds)
			seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Scopes(configuration.Paginate(page, pageSize)).Where("id in (?)", seriesIds).Where("series_type = ?", seriesType).Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("SeriesProvider.Provider").Preload("Episodes").Find(&seriesResult)
		}
	} else {
		if len(providers) == 0 {
			seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Scopes(configuration.Paginate(page, pageSize)).Where("series_type = ?", seriesType).Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("SeriesProvider.Provider").Preload("Episodes").Find(&seriesResult)
		} else {
			//var providerResult entity.Provider
			var providerIds []uint
			seriesRepository.DB.WithContext(ctx).Model(&entity.Provider{}).Where("hash_id in (?)", providers).Pluck("id", &providerIds)
			seriesRepository.DB.WithContext(ctx).Model(&entity.SeriesProvider{}).Where("provider_id in (?)", providerIds).Pluck("series_id", &seriesIds)
			seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Scopes(configuration.Paginate(page, pageSize)).Where("id in (?)", seriesIds).Where("series_type = ?", seriesType).Preload("Genres").Preload("Publishers").Preload("PublishDays").Preload("SeriesAuthors.Person").Preload("SeriesProvider.Provider").Preload("Episodes").Find(&seriesResult)
		}
	}

	var totalCount int64
	seriesRepository.DB.WithContext(ctx).Model(&entity.Series{}).Where("id in (?)", seriesIds).Count(&totalCount)

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
	}

	// paignation 정보 추가
	// totalCount, currentPage, nextPage, pageSize, totalPage entity
	// totalCount: 전체 데이터 수
	// currentPage: 현재 페이지
	// hasNext: 다음 페이지가 있는지 여부
	// pageSize: 페이지당 데이터 수
	// totalPage: 전체 페이지 수
	totalPage := (int(totalCount) / pageSize) + 1
	hasNext := func() bool {
		if page >= totalPage {
			return false
		}
		return true
	}()

	SeriesWithPagination := model.SeriesWithPagination{
		Series: seriesResult,
		Pagination: model.Pagination{
			TotalCount:  int(totalCount),
			CurrentPage: page,
			HasNext:     hasNext,
			PageSize:    pageSize,
			TotalPage:   totalPage,
		},
	}

	return SeriesWithPagination, nil
}

// Has Like Series
func (seriesRepository *seriesRepositoryImpl) HasLikeSeries(ctx context.Context, userId uint, seriesId uint) (bool, error) {
	var userLikeSeries entity.UserLikeSeries
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.UserLikeSeries{}).Where("user_id = ? and series_id = ?", userId, seriesId).First(&userLikeSeries)
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

// Like Series
func (seriesRepository *seriesRepositoryImpl) LikeSeries(ctx context.Context, userId uint, seriesId uint) error {
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.UserLikeSeries{}).Create(&entity.UserLikeSeries{UserId: userId, SeriesId: seriesId})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Unlike Series
func (seriesRepository *seriesRepositoryImpl) UnlikeSeries(ctx context.Context, userId uint, seriesId uint) error {
	var userLikeSeries entity.UserLikeSeries
	seriesRepository.DB.WithContext(ctx).Model(&entity.UserLikeSeries{}).Where("user_id = ? and series_id = ?", userId, seriesId).First(&userLikeSeries)
	result := seriesRepository.DB.WithContext(ctx).Model(&entity.UserLikeSeries{}).Delete(&userLikeSeries)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
