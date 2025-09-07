package repository

import (
	"context"
	"fmt"

	"github.com/confteam/bots-info-service/internal/dto"
	"github.com/confteam/bots-info-service/internal/logger"
	"github.com/confteam/bots-info-service/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Bot struct {
	ID         int
	Token      string
	Confession string
	ChatID     string
	ChannelID  string
	Type       string
	Code       string
}

type BotRepository struct {
	db *pgxpool.Pool
}

func NewBotRepository(db *pgxpool.Pool) *BotRepository {
	return &BotRepository{db: db}
}

func (r *BotRepository) RegisterBot(ctx context.Context, dto dto.RegisterDto) (*Bot, error) {
	code, err := utils.GenerateCode()
	if err != nil {
		return nil, fmt.Errorf("failed to generate code: %v", err)
	}

	bot := &Bot{
		Token: dto.Token,
		Type:  dto.Type,
		Code:  code,
	}

	query := "INSERT INTO bots (token, type, code) VALUES ($1, $2, $3) RETURNING id"
	err = r.db.QueryRow(ctx, query, bot.Token, bot.Type, bot.Code).Scan(&bot.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert bot: %v", err)
	}

	logger.Log.Info("registered bot", "type", dto.Type)

	return bot, nil
}
