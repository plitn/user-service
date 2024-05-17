package user_service

import (
	"context"
	"github.com/user-service/internal/model"
)

type Service interface {
	SaveUserInfo(ctx context.Context, user model.User) error
	GetUserInfo(ctx context.Context, userId int64) (model.User, error)
	UpdateUserInfo(ctx context.Context, user model.User) error
	GetUsersById(ctx context.Context, userIds []int64) ([]model.User, error)
	SaveStylistInfo(ctx context.Context, stylistInfo model.Stylist) error
	GetStylistById(ctx context.Context, stylistId int64) (model.Stylist, error)
	GetStylistsByIds(ctx context.Context, stylistId []int64) ([]model.Stylist, error)
	AddUserStylist(ctx context.Context, userId int64, stylistId int64) error
	DeleteUserStylist(ctx context.Context, userId int64, stylistId int64) error
	GetUserStylistsList(ctx context.Context, userId int64) ([]model.Stylist, error)
	GetStylistUsersList(ctx context.Context, stylistId int64) ([]model.User, error)

	CheckUserByName(ctx context.Context, username, password string) (model.User, error)
}
