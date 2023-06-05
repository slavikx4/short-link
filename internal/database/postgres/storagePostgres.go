package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/slavikx4/short-link/internal/models"
)

type StoragePostgres struct {
	DB *pgxpool.Pool
}

func NewStoragePostgres(db *pgxpool.Pool) *StoragePostgres {
	return &StoragePostgres{DB: db}
}

func (s *StoragePostgres) GetLinkFromStorage(shortLink string) (*models.Link, error) {

	link := models.Link{
		OriginalLink: "",
		ShortLink:    shortLink,
	}

	query := `SELECT "originalLink" FROM "Link" WHERE "shortLink"=$1`

	if err := s.DB.QueryRow(context.Background(), query, shortLink).Scan(&link.OriginalLink); err != nil {
		if err == pgx.ErrNoRows {
			return nil, err
			//TODO назначить ошибки в сервисе
		}
		return nil, err
	}

	return &link, nil
}

func (s *StoragePostgres) SetLinkIntoStorage(link *models.Link) error {
	query := `INSERT INTO "Link" ("shortLink", "originalLink") VALUES ($1, $2)`

	if _, err := s.DB.Exec(context.Background(), query, link.ShortLink, link.OriginalLink); err != nil {
		return err
	}

	return nil
}
