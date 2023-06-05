package inMemory

import (
	"errors"
	"github.com/slavikx4/short-link/internal/models"
)

type StorageInMemory struct {
	ShortToLongLink map[string]string
}

func NewStorageInMemory() *StorageInMemory {
	return &StorageInMemory{ShortToLongLink: make(map[string]string)}
}
func (s *StorageInMemory) GetLinkFromStorage(shortLink string) (*models.Link, error) {
	//по ключу shortLink пытаемся взять originalLink
	originalLink, ok := s.ShortToLongLink[shortLink]
	if ok == false {
		//если ссылки не было в базе, то нужно вернуть ошибку,
		//но так как у нас булевое ok, то возвращаем пустую ошибку
		return nil, errors.New("")
	}

	return &models.Link{ShortLink: shortLink, OriginalLink: originalLink}, nil
}

func (s *StorageInMemory) SetLinkIntoStorage(link *models.Link) error {

	//добавляем ссылку в хранилище, где ключ короткая ссылка
	s.ShortToLongLink[link.ShortLink] = link.OriginalLink
	//так как возникновенуть ошибка не может
	//но нужно реализовывать интерфейс, то возвращаем "заглушку"
	return nil

}
