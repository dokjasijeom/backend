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
	for i := 1; i <= int(toEpisodeNumber); i++ {
		episode := entity.Episode{
			EpisodeNumber: uint(i),
			Series:        []entity.Series{{Id: seriesId}},
		}
		episodes = append(episodes, episode)
	}
	result := episodeRepository.DB.WithContext(ctx).Model(&entity.Episode{}).Create(&episodes)
	if result.Error != nil {
		return nil, result.Error
	}

	return episodes, nil
}
