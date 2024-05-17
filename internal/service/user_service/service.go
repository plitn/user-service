package user_service

import (
	"context"
	"fmt"
	"github.com/user-service/internal/model"
	"github.com/user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) SaveUserInfo(ctx context.Context, user model.User) error {
	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("cannot hash password")
		}
		user.Password = string(hash)
	}
	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("cannot save user data: %w", err)
	}
	return nil
}

func (s *service) UpdateUserInfo(ctx context.Context, user model.User) error {
	err := s.repo.UpdateUserData(ctx, user)
	if err != nil {
		return fmt.Errorf("cannot update user data: %w", err)
	}
	return nil
}

func (s *service) GetUserInfo(ctx context.Context, userId int64) (model.User, error) {
	user, err := s.repo.GetUserByID(ctx, userId)
	if err != nil {
		return model.User{}, fmt.Errorf("cannot get user data: %w", err)
	}
	return user, nil
}

func (s *service) GetUsersById(ctx context.Context, userIds []int64) ([]model.User, error) {
	users, err := s.repo.GetUsersByIDs(ctx, userIds)
	if err != nil {
		return nil, fmt.Errorf("cannot get users data: %w", err)
	}
	return users, nil
}

func (s *service) SaveStylistInfo(ctx context.Context, stylistInfo model.Stylist) error {
	err := s.repo.CreateStylists(ctx, stylistInfo)
	if err != nil {
		return fmt.Errorf("cannot save stylist data: %w", err)
	}
	return nil
}

func (s *service) GetStylistById(ctx context.Context, stylistId int64) (model.Stylist, error) {
	stylist, err := s.repo.GetStylistByID(ctx, stylistId)
	if err != nil {
		return model.Stylist{}, fmt.Errorf("cannot get stylist data: %w", err)
	}
	return stylist, nil
}

func (s *service) GetStylistsByIds(ctx context.Context, stylistId []int64) ([]model.Stylist, error) {
	stylist, err := s.repo.GetStylistsByIDs(ctx, stylistId)
	if err != nil {
		return nil, fmt.Errorf("cannot get stylist data: %w", err)
	}
	return stylist, nil
}

func (s *service) AddUserStylist(ctx context.Context, userId int64, stylistId int64) error {
	pair := model.UserStylist{
		UserId:    userId,
		StylistId: stylistId,
	}
	err := s.repo.AddUserStylist(ctx, pair)
	if err != nil {
		return fmt.Errorf("cannot add stylist data: %w", err)
	}
	return nil
}

func (s *service) DeleteUserStylist(ctx context.Context, userId int64, stylistId int64) error {
	pair := model.UserStylist{
		UserId:    userId,
		StylistId: stylistId,
	}
	err := s.repo.RemoveUserStylist(ctx, pair)
	if err != nil {
		return fmt.Errorf("cannot delete stylist data: %w", err)
	}
	return nil
}

func (s *service) GetUserStylistsList(ctx context.Context, userId int64) ([]model.Stylist, error) {
	pairs, err := s.repo.GetPairsByUser(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("cannot get stylist data: %w", err)
	}
	stylists, err := s.repo.GetStylistsByIDs(ctx, pairs)
	if err != nil {
		return nil, fmt.Errorf("cannot get stylist data: %w", err)
	}
	return stylists, nil
}

func (s *service) GetStylistUsersList(ctx context.Context, stylistId int64) ([]model.User, error) {
	pairs, err := s.repo.GetPairsByStylist(ctx, stylistId)
	if err != nil {
		return nil, fmt.Errorf("cannot get users data: %w", err)
	}
	users, err := s.repo.GetUsersByIDs(ctx, pairs)
	if err != nil {
		return nil, fmt.Errorf("cannot get users data: %w", err)
	}
	return users, nil
}

func (s *service) CheckUserByName(ctx context.Context, username, password string) (model.User, error) {
	user, err := s.repo.GetUserCreditsByName(ctx, username)
	if err != nil {
		return model.User{}, fmt.Errorf("cannot get user data: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, fmt.Errorf("error matching password and hash, %v", err)
	}
	return user, nil
}
