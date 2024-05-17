package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/user-service/internal/model"
	"github.com/user-service/internal/service/user_service"
	"github.com/user-service/internal/service/ws"
	"log"
	"net/http"
	"strconv"
)

type service struct {
	userService       user_service.Service
	wsService         ws.Service
	upgrader          *websocket.Upgrader
	clientConnections map[string]map[*websocket.Conn]bool
	broadcast         chan model.Message
}

func NewHandler(userService user_service.Service, wsService ws.Service, upgrader websocket.Upgrader) *service {
	var clientConnections = make(map[string]map[*websocket.Conn]bool)
	var broadcast = make(chan model.Message)
	return &service{userService: userService, wsService: wsService, upgrader: &upgrader,
		clientConnections: clientConnections, broadcast: broadcast}
}

func (s *service) SaveUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req model.User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.userService.SaveUserInfo(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *service) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req model.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.userService.UpdateUserInfo(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *service) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userName := r.URL.Query().Get("name")
	if userName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userPassword := r.URL.Query().Get("password")
	if userPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := s.userService.CheckUserByName(ctx, userName, userPassword)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *service) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userId int64
	userIdString := r.URL.Query().Get("id")
	if userIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		userId, _ = strconv.ParseInt(userIdString, 10, 64)
	}

	f, err := s.userService.GetUserInfo(ctx, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *service) GetUsersByIds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ids, err := s.parseIdsArray(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := s.userService.GetUsersById(ctx, ids)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *service) GetStylistsByIds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ids, err := s.parseIdsArray(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := s.userService.GetStylistsByIds(ctx, ids)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *service) parseIdsArray(r *http.Request) ([]int64, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	var ids []int64

	idsString, ok := r.Form["id"]
	if !ok {
		return nil, err
	}

	for _, id := range idsString {
		newId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, newId)
	}
	return ids, nil
}

func (s *service) SaveStylist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req model.Stylist

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.userService.SaveStylistInfo(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *service) GetStylistById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var stylistId int64
	stylistIdString := r.URL.Query().Get("id")
	if stylistIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		stylistId, _ = strconv.ParseInt(stylistIdString, 10, 64)
	}

	f, err := s.userService.GetStylistById(ctx, stylistId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *service) GetUsersStylists(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userId int64
	userIdString := r.URL.Query().Get("id")
	if userIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		userId, _ = strconv.ParseInt(userIdString, 10, 64)
	}

	f, err := s.userService.GetUserStylistsList(ctx, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *service) GetStylistsUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var stylistId int64
	stylistIdString := r.URL.Query().Get("id")
	if stylistIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		stylistId, _ = strconv.ParseInt(stylistIdString, 10, 64)
	}

	f, err := s.userService.GetStylistUsersList(ctx, stylistId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *service) AddUserStylist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req model.UserStylistRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.userService.AddUserStylist(ctx, req.UserId, req.StylistId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *service) RemoveUserStylist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req model.UserStylistRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.userService.DeleteUserStylist(ctx, req.UserId, req.StylistId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *service) CreateRoom(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var stylistId, userId int64
	stylistIdString := r.URL.Query().Get("stylist_id")
	if stylistIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		stylistId, _ = strconv.ParseInt(stylistIdString, 10, 64)
	}

	userIdString := r.URL.Query().Get("user_id")
	if userIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		userId, _ = strconv.ParseInt(userIdString, 10, 64)
	}

	err := s.wsService.CreateRoom(ctx, userId, stylistId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *service) SaveMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req model.Message

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.wsService.SaveMessage(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *service) GetUserRooms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userId int64
	userIdString := r.URL.Query().Get("user_id")
	if userIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		userId, _ = strconv.ParseInt(userIdString, 10, 64)
	}

	rooms, err := s.wsService.GetUserRooms(ctx, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *service) GetStylistRooms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var stylistId int64
	stylistIdString := r.URL.Query().Get("stylist_id")
	if stylistIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		stylistId, _ = strconv.ParseInt(stylistIdString, 10, 64)
	}

	rooms, err := s.wsService.GetStylistRooms(ctx, stylistId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rooms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *service) GetChatMessages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var chatId int64
	chatIdString := r.URL.Query().Get("chat_id")
	if chatIdString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		chatId, _ = strconv.ParseInt(chatIdString, 10, 64)
	}

	msgs, err := s.wsService.GetChatMessages(ctx, chatId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(msgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *service) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error upgrading HTTP connection: %v\n", err)
	}
	defer ws.Close()

	username := r.URL.Query().Get("user_id")
	chatID := r.URL.Query().Get("chat_id")

	if chatID == "" || username == "" {
		fmt.Println("Missing username or chat_id")
		return
	}

	chatIdInt, _ := strconv.ParseInt(chatID, 10, 64)

	if _, ok := s.clientConnections[chatID]; !ok {
		s.clientConnections[chatID] = make(map[*websocket.Conn]bool)
	}
	s.clientConnections[chatID][ws] = true

	dbChat := &model.ChatUsers{ChatId: chatIdInt}

	for {
		var msg model.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading JSON: %v", err)
			delete(s.clientConnections[chatID], ws)
			break
		}
		msg.ChatID = dbChat.ChatId
		err = s.wsService.SaveMessage(context.Background(), msg)
		if err != nil {
			log.Printf("Error saving message: %v", err)
			return
		}
		s.broadcast <- msg
	}
}

func (s *service) HandleMessages() {
	for {
		msg := <-s.broadcast
		for client := range s.clientConnections[strconv.Itoa(int(msg.ChatID))] {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error writing JSON: %v", err)
				client.Close()
				delete(s.clientConnections[strconv.Itoa(int(msg.ChatID))], client)
			}
		}
	}
}
