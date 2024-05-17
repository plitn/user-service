package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/user-service/internal/repository"
	"github.com/user-service/internal/service/handler"
	"github.com/user-service/internal/service/user_service"
	"github.com/user-service/internal/service/ws"
	"net/http"
	"os"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	connStr := fmt.Sprintf("host=%s port=5432 user=postgres password=postgres dbname=postgres sslmode=disable timezone=UTC", dbHost)

	dbInst, err := sqlx.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("failed to connect database, err: %v", err)
		return
	}
	err = dbInst.Ping()
	if err != nil {
		fmt.Printf("failed to ping database, err: %v", err)
		return
	}
	defer dbInst.Close()
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	repo := repository.NewRepository(dbInst)
	userService := user_service.NewUserService(repo)
	wsService := ws.NewWsService(repo)
	handlerService := handler.NewHandler(userService, wsService, upgrader)

	mux := chi.NewRouter()
	mux.Post("/save-user", handlerService.SaveUser)
	mux.Get("/get-user", handlerService.GetUserById)
	mux.Get("/get-users", handlerService.GetUsersByIds)
	mux.Post("/save-stylist", handlerService.SaveStylist)
	mux.Get("/get-stylist", handlerService.GetStylistById)
	mux.Get("/get-stylists", handlerService.GetStylistsByIds)
	mux.Get("/get-user-stylists", handlerService.GetUsersStylists)
	mux.Get("/get-stylist-users", handlerService.GetStylistsUsers)
	mux.Post("/add-stylist", handlerService.AddUserStylist)
	mux.Delete("/remove-stylist", handlerService.RemoveUserStylist)
	mux.Put("/update-user", handlerService.UpdateUser)
	mux.Get("/login", handlerService.LoginUser)

	mux.Post("/create-room", handlerService.CreateRoom)
	mux.Post("/save-message", handlerService.SaveMessage)
	mux.Get("/get-user-rooms", handlerService.GetUserRooms)
	mux.Get("/get-stylist-rooms", handlerService.GetStylistRooms)
	mux.Get("/get-chat-messages", handlerService.GetChatMessages)

	mux.Get("/ws", handlerService.HandleWebSocket)

	go handlerService.HandleMessages()

	httpServer := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", 8083),
		Handler: mux,
	}

	fmt.Printf("listening to http://0.0.0.0:%d/ for debug http", 8083)
	if err = httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("failed to listen on port 8080: %v", err)
	}
}
