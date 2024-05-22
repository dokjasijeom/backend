package impl

import (
	"context"
	"github.com/dokjasijeom/backend/entity"
	"github.com/dokjasijeom/backend/repository"
	"github.com/dokjasijeom/backend/service"
	"github.com/samber/lo"
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
	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.CreateUserRecordSeriesEpisode(ctx, userRecordSeriesEpisode)
}

// Create bulk user record series episode
func (userRecordSeriesEpisodeService *userRecordSeriesEpisodeServiceImpl) CreateBulkUserRecordSeriesEpisode(ctx context.Context, userRecordSeriesEpisodes []entity.UserRecordSeriesEpisode) ([]entity.UserRecordSeriesEpisode, error) {
	currentHasUserRecordSeriesEpisodes, err := userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.GetUserRecordSeriesEpisodeByUserRecordSeriesId(ctx, userRecordSeriesEpisodes[0].UserRecordSeriesId)
	if err != nil {
		return nil, err
	}

	clonedUserRecordSeriesEpisodes := userRecordSeriesEpisodes
	// 만약 다중 회차를 기록하려는 중 이미 DB에 기록된 회차가 있다면 해당 회차를 제외한다.
	if len(currentHasUserRecordSeriesEpisodes) > 0 {
		// for문 반복을 통해 존재하는지 체크
		for _, userRecordSeriesEpisode := range userRecordSeriesEpisodes {
			_, ok := lo.Find(currentHasUserRecordSeriesEpisodes, func(item entity.UserRecordSeriesEpisode) bool {
				return item.EpisodeId == userRecordSeriesEpisode.EpisodeId
			})
			_, index, isFind := lo.FindIndexOf(clonedUserRecordSeriesEpisodes, func(item entity.UserRecordSeriesEpisode) bool {
				return item.EpisodeId == userRecordSeriesEpisode.EpisodeId
			})

			if ok && isFind {
				// 존재한다면 해당 회차를 제외한다.
				clonedUserRecordSeriesEpisodes = append(clonedUserRecordSeriesEpisodes[:index], clonedUserRecordSeriesEpisodes[index+1:]...)
			}
		}
	}

	return userRecordSeriesEpisodeService.UserRecordSeriesEpisodeRepository.CreateBulkUserRecordSeriesEpisode(ctx, clonedUserRecordSeriesEpisodes)
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
