package service

import (
	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

type UrlService interface {
	FindAll() ([]model.Url, error)
	FindOne(urlId int) (*model.Url, error)
	Create(userId int, createUrlRequest dto.CreateUrlRequest) error
	Delete(userId, urlId int) error
}
