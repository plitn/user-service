package handler

import "net/http"

type Service interface {
	SaveUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	GetUsersByIds(w http.ResponseWriter, r *http.Request)
	GetStylistsByIds(w http.ResponseWriter, r *http.Request)
	SaveStylist(w http.ResponseWriter, r *http.Request)
	GetStylistById(w http.ResponseWriter, r *http.Request)
	GetUsersStylists(w http.ResponseWriter, r *http.Request)
	GetStylistsUsers(w http.ResponseWriter, r *http.Request)
	AddUserStylist(w http.ResponseWriter, r *http.Request)
	RemoveUserStylist(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)

	CreateRoom(w http.ResponseWriter, r *http.Request)
	SaveMessage(w http.ResponseWriter, r *http.Request)
	GetUserRooms(w http.ResponseWriter, r *http.Request)
	GetStylistRooms(w http.ResponseWriter, r *http.Request)
	GetChatMessages(w http.ResponseWriter, r *http.Request)
}
