package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"gorm.io/gorm"
)

func NewEpisodeRepositoryImpl(DB *gorm.DB) repository.EpisodeRepository {
	return &episodeRepositoryImpl{DB: DB}
}

type episodeRepositoryImpl struct {
	*gorm.DB
}

// Create New Episode
func (episodeRepository *episodeRepositoryImpl) CreateEpisode(ctx context.Context, episodeNumber uint) (entity.Episode, error) {
	episode := entity.Episode{
		EpisodeNumber: episodeNumber,
	}
	result := episodeRepository.DB.WithContext(ctx).Model(&entity.Episode{}).Create(&episode)
	if result.Error != nil {
		return entity.Episode{}, result.Error
	}
	return episode, nil
}

// Get All Episode
func (episodeRepository *episodeRepositoryImpl) GetAllEpisode(ctx context.Context) ([]entity.Episode, error) {
	var episodeResult []entity.Episode
	result := episodeRepository.DB.WithContext(ctx).Model(&entity.Episode{}).Find(&episodeResult)
	if result.Error != nil {
		return nil, result.Error
	}
	return episodeResult, nil
}

// Delete Episode
func (episodeRepository *episodeRepositoryImpl) DeleteEpisode(ctx context.Context, episodeId uint) error {
	result := episodeRepository.DB.WithContext(ctx).Model(&entity.Episode{}).Where("id = ?", episodeId).Delete(&entity.Episode{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update Episode
func (episodeRepository *episodeRepositoryImpl) UpdateEpisode(ctx context.Context, episodeId uint, episodeNumber uint) error {
	result := episodeRepository.DB.WithContext(ctx).Model(&entity.Episode{}).Where("id = ?", episodeId).Updates(entity.Episode{
		EpisodeNumber: episodeNumber,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get Episode By Id
func (episodeRepository *episodeRepositoryImpl) GetEpisodeById(ctx context.Context, episodeId uint) (entity.Episode, error) {
	var episodeResult entity.Episode
	result := episodeRepository.DB.WithContext(ctx).Model(&entity.Episode{}).Where("id = ?", episodeId).Find(&episodeResult)
	if result.Error != nil {
		return entity.Episode{}, result.Error
	}
	return episodeResult, nil
}

// Create Bulk New Episode
func (episodeRepository *episodeRepositoryImpl) CreateBulkEpisode(ctx context.Context, seriesId, toEpisodeNumber uint) ([]entity.Episode, error) {
	var episodes []entity.Episode
	var currentSeries entity.Series

	startNumber := 1

	var lastEpisode entity.Episode
	result := episodeRepository.DB.WithContext(ctx).Model(&entity.Series{}).Where("id = ?", seriesId).Preload("Episodes", func(db *gorm.DB) *gorm.DB {
		return db.Order("episode_number desc")
	}).First(&currentSeries)

	// seriesId에 해당하는 Series가 존재하지 않으면 에러 반환
	if result.Error != nil {
		return nil, result.Error
	}

	// Series에 Episode가 존재하지 않으면 startNumber = 1
	if len(currentSeries.Episodes) == 0 {
		startNumber = 1
	} else {
		lastEpisode = currentSeries.Episodes[0]
	}

	// 요청 받은 toEpisodeNumber가 이미 존재하는 Episode의 EpisodeNumber보다 작으면 에러 반환
	if lastEpisode.EpisodeNumber >= toEpisodeNumber {
		return nil, gorm.ErrRecordNotFound
	}

	// lastEpisode가 존재하면 lastEpisode.EpisodeNumber + 1부터 toEpisodeNumber까지 생성
	if lastEpisode.EpisodeNumber != 0 {
		startNumber = int(lastEpisode.EpisodeNumber) + 1
	}

	for i := startNumber; i <= int(toEpisodeNumber); i++ {
		episode := entity.Episode{
			EpisodeNumber: uint(i),
			Series:        []entity.Series{{Id: seriesId}},
		}
		episodes = append(episodes, episode)
	}
	result = episodeRepository.DB.WithContext(ctx).Model(&entity.Episode{}).Create(&episodes)
	if result.Error != nil {
		return nil, result.Error
	}

	return episodes, nil
}
