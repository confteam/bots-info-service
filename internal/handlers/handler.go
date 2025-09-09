package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/confteam/bots-info-service/internal/config"
	"github.com/confteam/bots-info-service/internal/logger"
	"github.com/confteam/bots-info-service/internal/repository"
)

type Handler struct {
	repo   *repository.BotRepository
	config config.Config
}

func Start(repo *repository.BotRepository) error {
	config := config.GetConfig()
	r := http.NewServeMux()
	h := Handler{
		repo:   repo,
		config: config,
	}

	r.HandleFunc("POST /register", h.Register)

	logger.Log.Info("server started", "port", config.Port)
	return http.ListenAndServe(":"+config.Port, r)
}

func (h Handler) CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	return ctx, cancel
}
