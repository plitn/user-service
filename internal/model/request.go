package model

type UserStylistRequest struct {
	UserId    int64 `json:"user_id"`
	StylistId int64 `json:"stylist_id"`
}
