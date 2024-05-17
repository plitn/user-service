package repository

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/user-service/internal/model"
)

const (
	usersTable          = "users"
	stylistsTable       = "stylists"
	stylistsSkillsTable = "stylists_skills"
	usersStylistsTable  = "users_stylists"
	loginTable          = "users_login"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
func (r *repository) CreateUser(ctx context.Context, user model.User) error {
	query := `insert into users (name, email, password, image_url, gender, age, weight) values ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.ImageUrl, user.Gender, user.Age, user.Weight)
	if err != nil {
		return fmt.Errorf("cannot save user: %w", err)
	}
	return nil
}

func (r *repository) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	var user model.User
	query, _, err := goqu.From(usersTable).Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return model.User{}, fmt.Errorf("cannot configure query: %w", err)
	}
	err = r.db.GetContext(ctx, &user, query)
	if err != nil {
		return model.User{}, fmt.Errorf("cannot save user: %w", err)
	}
	return user, nil
}

func (r *repository) GetUsersByIDs(ctx context.Context, id []int64) ([]model.User, error) {
	var users []model.User
	query, _, err := goqu.From(usersTable).Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("cannot configure query: %w", err)
	}
	err = r.db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, fmt.Errorf("cannot save user: %w", err)
	}
	return users, nil
}

func (r *repository) CreateStylists(ctx context.Context, stylist model.Stylist) error {
	query := `insert into stylists (name, email, password,  image_url, gender, age, experience_time, portfolio, experience_description) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.ExecContext(ctx, query, stylist.Name, stylist.Email, stylist.Password, stylist.ImageUrl, stylist.Gender,
		stylist.Age, stylist.ExperienceTime, stylist.Portfolio, stylist.ExperienceDescription)
	if err != nil {
		return fmt.Errorf("cannot save stylist: %w", err)
	}
	return nil
}

func (r *repository) GetStylistByID(ctx context.Context, id int64) (model.Stylist, error) {
	var stylist model.Stylist
	query, _, err := goqu.From(stylistsTable).Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return model.Stylist{}, fmt.Errorf("cannot configure query: %w", err)
	}
	err = r.db.GetContext(ctx, &stylist, query)
	if err != nil {
		return model.Stylist{}, fmt.Errorf("cannot save user: %w", err)
	}
	return stylist, nil
}

func (r *repository) GetStylistsByIDs(ctx context.Context, id []int64) ([]model.Stylist, error) {
	var stylists []model.Stylist
	query, _, err := goqu.From(stylistsTable).Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("cannot configure query: %w", err)
	}
	err = r.db.SelectContext(ctx, &stylists, query)
	if err != nil {
		return nil, fmt.Errorf("cannot save user: %w", err)
	}
	return stylists, nil
}

func (r *repository) AddUserStylist(ctx context.Context, pair model.UserStylist) error {
	query := `INSERT INTO users_stylists (user_id, stylist_id) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, pair.UserId, pair.StylistId)
	if err != nil {
		return fmt.Errorf("cannot add stylist to user: %w", err)
	}
	return nil
}

func (r *repository) RemoveUserStylist(ctx context.Context, pair model.UserStylist) error {
	query := `DELETE FROM users_stylists WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, pair.UserId)
	if err != nil {
		return fmt.Errorf("cannot remove stylist from user: %w", err)
	}
	return nil
}

func (r *repository) GetPairsByUser(ctx context.Context, userId int64) ([]int64, error) {
	var pairs []int64
	query, _, err := goqu.From(usersStylistsTable).Select("stylist_id").Where(goqu.Ex{"user_id": userId}).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("cannot configure query: %w", err)
	}
	err = r.db.SelectContext(ctx, &pairs, query)
	if err != nil {
		return nil, fmt.Errorf("cannot save user: %w", err)
	}
	return pairs, nil
}

func (r *repository) GetPairsByStylist(ctx context.Context, stylistId int64) ([]int64, error) {
	var pairs []int64
	query, _, err := goqu.From(usersStylistsTable).Select("user_id").Where(goqu.Ex{"stylist_id": stylistId}).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("cannot configure query: %w", err)
	}
	err = r.db.SelectContext(ctx, &pairs, query)
	if err != nil {
		return nil, fmt.Errorf("cannot save user: %w", err)
	}
	return pairs, nil
}

func (r *repository) GetUserCreditsByName(ctx context.Context, username string) (model.User, error) {
	var user model.User
	query, _, err := goqu.From(usersTable).Select().Where(goqu.Ex{"name": username}).ToSQL()
	if err != nil {
		return model.User{}, fmt.Errorf("cannot configure query: %w", err)
	}
	err = r.db.GetContext(ctx, &user, query)
	if err != nil {
		return model.User{}, fmt.Errorf("cannot get user credits: %w", err)
	}
	return user, nil
}

func (r *repository) UpdateUserData(ctx context.Context, user model.User) error {
	updateMap := map[string]interface{}{}
	if user.Name != "" {
		updateMap["name"] = user.Name
	}
	if user.Email != "" {
		updateMap["email"] = user.Email
	}
	if user.Password != "" {
		updateMap["password"] = user.Password
	}
	if user.ImageUrl != "" {
		updateMap["image_url"] = user.ImageUrl
	}
	if user.Gender != 0 {
		updateMap["gender"] = user.Gender
	}
	if user.Age != 0 {
		updateMap["age"] = user.Age
	}
	if user.Weight != 0 {
		updateMap["weight"] = user.Weight
	}
	query, _, err := goqu.Update(usersTable).Set(updateMap).Where(goqu.Ex{"id": user.Id}).ToSQL()
	if err != nil {
		return fmt.Errorf("cannot configure query: %w", err)
	}
	_, err = r.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("cannot update user: %w", err)
	}
	return nil
}

func (r *repository) SaveMessage(ctx context.Context, message model.Message) error {
	query := `INSERT INTO messages (chat_id, user_id, text) values ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, message.ChatID, message.UserId, message.Text)
	if err != nil {
		return fmt.Errorf("cannot save message: %w", err)
	}
	return nil
}

func (r *repository) CreateRoom(ctx context.Context, userId, stylistId int64) error {
	query := `INSERT INTO rooms (user_id, stylist_id) values ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, userId, stylistId)
	if err != nil {
		return fmt.Errorf("cannot create room: %w", err)
	}
	return nil
}

func (r *repository) GetUserRooms(ctx context.Context, userId int64) ([]model.ChatUsers, error) {
	var chatUsersData []model.ChatUsers
	query, _, err := goqu.From("rooms").Where(goqu.Ex{"user_id": userId}).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("cannot configure query: %w", err)
	}
	err = r.db.SelectContext(ctx, &chatUsersData, query)
	if err != nil {
		return nil, fmt.Errorf("cannot get rooms: %w", err)
	}
	return chatUsersData, nil
}

func (r *repository) GetStylistRooms(ctx context.Context, stylist int64) ([]model.ChatUsers, error) {
	var chatUsersData []model.ChatUsers
	query, _, err := goqu.From("rooms").Where(goqu.Ex{"stylist_id": stylist}).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("cannot configure query: %w", err)
	}
	err = r.db.SelectContext(ctx, &chatUsersData, query)
	if err != nil {
		return nil, fmt.Errorf("cannot get rooms: %w", err)
	}
	return chatUsersData, nil
}

func (r *repository) GetChatMessages(ctx context.Context, chatId int64) ([]model.Message, error) {
	var chatMessagesData []model.Message
	query, _, err := goqu.From("messages").Where(goqu.Ex{"chat_id": chatId}).Order(goqu.C("id").Asc()).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("cannot configure query: %w", err)
	}
	err = r.db.SelectContext(ctx, &chatMessagesData, query)
	if err != nil {
		return nil, fmt.Errorf("cannot get messages: %w", err)
	}
	return chatMessagesData, nil
}
