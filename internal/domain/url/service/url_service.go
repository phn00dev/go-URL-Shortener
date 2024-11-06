package service

import (
	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

type UrlService interface {
	FindAll() ([]model.Url, error)
	FindOne(urlId int) (*model.Url, error)
	Create(userId int, createUrlRequest dto.CreateUrlRequest) (*model.Url, error)
	Delete(userId, urlId int) error

	// user urls
	FindAllUserUrls(userId int) ([]model.Url, error)
	FindOneUserUrl(userId, urlId int) (*model.Url, error)
	GetByShortUrl(shortUrl string) (*model.Url, error)
	UpdateClickCount(urlId, clickCount int) error
	SaveUrlAccessLog(accessLog model.UrlAccessLog) error
}
