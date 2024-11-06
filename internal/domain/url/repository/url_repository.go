package repository

import "github.com/phn00dev/go-URL-Shortener/internal/model"

type UrlRepository interface {
	GetUrlById(urlId int) (*model.Url, error)
	GetAllUserUrl(userId int) ([]model.Url, error)
	GetAllUrl() ([]model.Url, error)
	Create(url model.Url) error
	Delete(userId, urlId int) error
	GetUrlShortUrl(shortUrl string) (*model.Url, error)
}
