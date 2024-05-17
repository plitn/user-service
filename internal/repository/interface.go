package repository

import (
	"context"
	"github.com/user-service/internal/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByID(ctx context.Context, id int64) (model.User, error)
	GetUsersByIDs(ctx context.Context, id []int64) ([]model.User, error)
	CreateStylists(ctx context.Context, stylist model.Stylist) error
	GetStylistByID(ctx context.Context, id int64) (model.Stylist, error)
	GetStylistsByIDs(ctx context.Context, id []int64) ([]model.Stylist, error)
	AddUserStylist(ctx context.Context, pair model.UserStylist) error
	RemoveUserStylist(ctx context.Context, pair model.UserStylist) error
	GetPairsByUser(ctx context.Context, userId int64) ([]int64, error)
	GetPairsByStylist(ctx context.Context, stylistId int64) ([]int64, error)

	GetUserCreditsByName(ctx context.Context, username string) (model.User, error)

	UpdateUserData(ctx context.Context, user model.User) error

	SaveMessage(ctx context.Context, message model.Message) error
	CreateRoom(ctx context.Context, userId, stylistId int64) error
	GetUserRooms(ctx context.Context, userId int64) ([]model.ChatUsers, error)
	GetStylistRooms(ctx context.Context, stylist int64) ([]model.ChatUsers, error)
	GetChatMessages(ctx context.Context, chatId int64) ([]model.Message, error)
}
