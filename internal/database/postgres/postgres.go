package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/slavikx4/short-link/pkg/logger"
)

func NewPoolPostgres(config string) (*pgxpool.Pool, error) {
	url := config
	DB, err := pgxpool.New(context.Background(), url)
	if err != nil {
		logger.Logger.Error.Println("не удалось подключиться к DataBase Link: ", err)
		return nil, err
	}
	if err := DB.Ping(context.Background()); err != nil {
		logger.Logger.Error.Println("не удалось пингануть к DataBase Link: ", err)
		return nil, err
	}
	logger.Logger.Process.Println("подключён успешно postgres")

	return DB, nil
}
