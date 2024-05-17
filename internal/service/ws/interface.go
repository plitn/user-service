package ws

import (
	"context"
	"github.com/user-service/internal/model"
)

type Service interface {
	SaveMessage(ctx context.Context, msg model.Message) error
	CreateRoom(ctx context.Context, userId, stylistId int64) error
	GetUserRooms(ctx context.Context, userId int64) ([]model.ChatUsers, error)
	GetStylistRooms(ctx context.Context, stylistId int64) ([]model.ChatUsers, error)
	GetChatMessages(ctx context.Context, chatId int64) ([]model.Message, error)
}
