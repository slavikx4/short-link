package services

import (
	"github.com/slavikx4/short-link/internal/contracts"
	"github.com/slavikx4/short-link/internal/models"
	"github.com/slavikx4/short-link/pkg/hash"
	"github.com/slavikx4/short-link/pkg/logger"
	"unicode/utf8"
)

type Service struct {
	storage contracts.Storage
}

func NewService(storage contracts.Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) GetLink(shortLink string) (*models.Link, error) {

	if err := ValidateShortLink(shortLink); err != nil {
		return nil, err
	}

	link, err := s.storage.GetLinkFromStorage(shortLink)
	if err != nil {
		return nil, ErrorNoLink
	}

	return link, nil
}

func (s *Service) SetLink(originalLink string) (*models.Link, error) {

	//проверка на валидность вводимой ссылки
	if err := ValidateOriginalLink(originalLink); err != nil {
		return nil, err
	}

	//преобразуем originalLink в shortLink
	shortLink := hash.HashLink(originalLink)
	//так как хеш функция соответствует приципу детерменизма
	//то можно не вытаксивать из БД каждый раз shortLink,
	//а генерировать shortLink заново
	link := &models.Link{
		OriginalLink: originalLink,
		ShortLink:    shortLink,
	}

	if err := s.storage.SetLinkIntoStorage(link); err != nil {
		logger.Logger.Error.Println("ошибка добавления ссылки в бд: ", err)
		//ошибка возникает, если такая ссылка уже была сохранена
		//поэтому отправляется shortLink клиенту, игнорируя ошибку
		return link, nil
	}

	return link, nil
}

func ValidateOriginalLink(originalLink string) error {
	//TODO можно написать более сложную проверку, но не уверен, что это требуется от меня
	if originalLink == "" {
		return ErrorInvalidValue
	}
	return nil
}

func ValidateShortLink(shortLink string) error {
	if utf8.RuneCountInString(shortLink) != 10 {
		return ErrorInvalidValue
	}
	return nil
}
