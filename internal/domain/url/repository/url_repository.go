package repository

import "github.com/phn00dev/go-URL-Shortener/internal/model"

type UrlRepository interface {
	GetUrlById(userId, urlId int) (*model.Url, error)
	GetAllUserUrl(userId int) ([]model.Url, error)
	GetAllUrl() ([]model.Url, error)
	Create(userId int, url model.Url) error
	Update(urlId int, url model.Url) error
	Delete(userId, urlId int) error
}
