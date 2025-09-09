package db

import (
	"context"
	"fmt"
	"time"

	"github.com/confteam/bots-info-service/internal/config"
	"github.com/confteam/bots-info-service/internal/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB() (*pgxpool.Pool, error) {
	config := config.GetConfig()
	pgConn := fmt.Sprintf("postgres://%s:%s@postgres:5432/bots", config.PostgresUser, config.PostgresPassword)
	logger.Log.Info("created postgresql connection")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pool, err := pgxpool.New(ctx, pgConn)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to db: %v", err)
	}
	logger.Log.Info("connected to db")

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while checking connection to db: %v", err)
	}
	logger.Log.Info("checked connection")

	return pool, nil
}
