package model

type Chat struct {
	ChatId   int64
	Users    []*User
	Messages []Message
}

type Message struct {
	Id     int64  `json:"id" db:"id"`
	ChatID int64  `json:"chat_id" db:"chat_id"`
	UserId int64  `json:"user_id" db:"user_id"`
	Text   string `json:"text" db:"text"`
}

type ChatUsers struct {
	ChatId    int64 `json:"chat_id" db:"chat_id"`
	UserId    int64 `json:"user_id" db:"user_id"`
	StylistId int64 `json:"stylist_id" db:"stylist_id"`
}
