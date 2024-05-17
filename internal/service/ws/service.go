package ws

import (
	"context"
	"github.com/user-service/internal/model"
	"github.com/user-service/internal/repository"
)

type service struct {
	repo repository.Repository
}

func NewWsService(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) SaveMessage(ctx context.Context, msg model.Message) error {
	err := s.repo.SaveMessage(ctx, msg)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CreateRoom(ctx context.Context, userId, stylistId int64) error {
	err := s.repo.CreateRoom(ctx, userId, stylistId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetUserRooms(ctx context.Context, userId int64) ([]model.ChatUsers, error) {
	rooms, err := s.repo.GetUserRooms(ctx, userId)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s *service) GetStylistRooms(ctx context.Context, stylistId int64) ([]model.ChatUsers, error) {
	rooms, err := s.repo.GetStylistRooms(ctx, stylistId)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s *service) GetChatMessages(ctx context.Context, chatId int64) ([]model.Message, error) {
	messages, err := s.repo.GetChatMessages(ctx, chatId)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
