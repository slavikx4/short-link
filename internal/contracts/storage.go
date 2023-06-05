package contracts

import "github.com/slavikx4/short-link/internal/models"

type Storage interface {
	GetLinkFromStorage(shortLink string) (*models.Link, error)
	SetLinkIntoStorage(link *models.Link) error
}
