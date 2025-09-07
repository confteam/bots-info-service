package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/confteam/bots-info-service/internal/dto"
	"github.com/confteam/bots-info-service/internal/logger"
)

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("got request", "method", r.Method, "path", r.URL.Path)

	var req dto.RegisterDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := h.CreateContext()
	defer cancel()

	bot, err := h.repo.RegisterBot(ctx, req)
	if err != nil {
		http.Error(w, "Failed to register bot "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bot)
}
