package impl

import (
	"context"
	"errors"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
)

func NewUserRecordSeriesEpisodeServiceImpl(userRecordSeriesEpisodeRepository *repository.UserRecordSeriesEpisodeRepository) service.UserRecordSeriesEpisodeService {
	return &userRecordSeriesEpisodeServiceImpl{UserRecordSeriesEpisodeRepository: *userRecordSeriesEpisodeRepository}
}

type userRecordSeriesEpisodeServiceImpl struct {
	repository.UserRecordSeriesEpisodeRepository
}

// Get user record series episode by id
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) GetUserRecordSeriesEpisodeById(ctx context.Context, id uint) (entity.UserRecordSeriesEpisode, error) {
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.GetUserRecordSeriesEpisodeById(ctx, id)
}

// Create user record series episode
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) CreateUserRecordSeriesEpisode(ctx context.Context, userRecordSeriesEpisode entity.UserRecordSeriesEpisode) (entity.UserRecordSeriesEpisode, error) {
	currentHasUserRecordSeriesEpisodes, err := userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.GetUserRecordSeriesEpisodeByUserRecordSeriesId(ctx, userRecordSeriesEpisode.UserRecordSeriesId)
	if err != nil {
		return entity.UserRecordSeriesEpisode{}, err
	}
	if len(currentHasUserRecordSeriesEpisodes) > 0 {
		// for문 반복을 통해 존재하는지 체크
		for _, currentHasUserRecordSeriesEpisode := range currentHasUserRecordSeriesEpisodes {
			if userRecordSeriesEpisode.EpisodeId != 0 {
				if currentHasUserRecordSeriesEpisode.EpisodeId == userRecordSeriesEpisode.EpisodeId {
					return entity.UserRecordSeriesEpisode{}, errors.New("already exists recorded episode")
				}
			} else {
				if currentHasUserRecordSeriesEpisode.EpisodeNumber == userRecordSeriesEpisode.EpisodeNumber {
					return entity.UserRecordSeriesEpisode{}, errors.New("already exists recorded episode")
				}
			}
		}
	}

	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.CreateUserRecordSeriesEpisode(ctx, userRecordSeriesEpisode)
}

// Create bulk user record series episode
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) CreateBulkUserRecordSeriesEpisode(ctx context.Context, userRecordSeriesEpisodes []entity.UserRecordSeriesEpisode) ([]entity.UserRecordSeriesEpisode, error) {
	currentHasUserRecordSeriesEpisodes, err := userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.GetUserRecordSeriesEpisodeByUserRecordSeriesId(ctx, userRecordSeriesEpisodes[0].UserRecordSeriesId)
	if err != nil {
		return nil, err
	}

	var deleteRecordIds []uint
	// 만약 다중 회차를 기록하려는 중 이미 DB에 기록된 회차가 있다면 해당 회차를 삭제하도록 아이디를 저장한다.
	if len(currentHasUserRecordSeriesEpisodes) > 0 {
		// for문 반복을 통해 존재하는지 체크
		for i := 0; i < len(userRecordSeriesEpisodes); i++ {
			if userRecordSeriesEpisodes[i].EpisodeId != 0 {
				for _, currentHasUserRecordSeriesEpisode := range currentHasUserRecordSeriesEpisodes {
					if currentHasUserRecordSeriesEpisode.EpisodeId == userRecordSeriesEpisodes[i].EpisodeId {
						deleteRecordIds = append(deleteRecordIds, currentHasUserRecordSeriesEpisode.Id)
					}
				}
			} else {
				for _, currentHasUserRecordSeriesEpisode := range currentHasUserRecordSeriesEpisodes {
					if currentHasUserRecordSeriesEpisode.EpisodeNumber == userRecordSeriesEpisodes[i].EpisodeNumber {
						deleteRecordIds = append(deleteRecordIds, currentHasUserRecordSeriesEpisode.Id)
					}
				}
			}
		}
	}

	if len(deleteRecordIds) > 0 {
		err := userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.DeleteUserRecordSeriesEpisodeByUserRecordSeriesEpisodeIds(ctx, deleteRecordIds)
		if err != nil {
			return nil, err
		}
	}

	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.CreateBulkUserRecordSeriesEpisode(ctx, userRecordSeriesEpisodes)
}

// Update user record series episode by id
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) UpdateUserRecordSeriesEpisodeById(ctx context.Context, id uint, userRecordSeriesEpisode entity.UserRecordSeriesEpisode) (entity.UserRecordSeriesEpisode, error) {
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.UpdateUserRecordSeriesEpisodeById(ctx, id, userRecordSeriesEpisode)
}

// Delete user record series episode by id
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) DeleteUserRecordSeriesEpisodeById(ctx context.Context, id uint) error {
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.DeleteUserRecordSeriesEpisodeById(ctx, id)
}

// Delete user record series episode by user record series episode ids
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) DeleteUserRecordSeriesEpisodeByUserRecordSeriesEpisodeIds(ctx context.Context, userRecordSeriesEpisodeIds []uint) error {
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.DeleteUserRecordSeriesEpisodeByUserRecordSeriesEpisodeIds(ctx, userRecordSeriesEpisodeIds)
}

// Delete user record series episode by user record series id
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) DeleteUserRecordSeriesEpisodeByUserRecordSeriesId(ctx context.Context, userRecordSeriesId uint) error {
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.DeleteUserRecordSeriesEpisodeByUserRecordSeriesId(ctx, userRecordSeriesId)
}
