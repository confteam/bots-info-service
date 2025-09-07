package main

import (
	"github.com/confteam/bots-info-service/internal/db"
	"github.com/confteam/bots-info-service/internal/handlers"
	"github.com/confteam/bots-info-service/internal/logger"
	"github.com/confteam/bots-info-service/internal/repository"
)

func main() {
	conn, err := db.ConnectToDB()
	if err != nil {
		logger.Log.Error("failed to connect to db", "error", err)
		return
	}
	defer conn.Close()

	repo := repository.NewBotRepository(conn)
	if err := handlers.Start(repo); err != nil {
		logger.Log.Error("failed to start handlers", "error", err)
		return
	}
}
