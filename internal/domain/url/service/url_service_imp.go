package service

import (
	"errors"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/repository"
	userRepository "github.com/phn00dev/go-URL-Shortener/internal/domain/user/repository"
	"github.com/phn00dev/go-URL-Shortener/internal/model"

)

type urlServiceImp struct {
	urlRepo  repository.UrlRepository
	userRepo userRepository.UserRepository
}

func NewUrlService(urlRepo repository.UrlRepository, userRepo userRepository.UserRepository) UrlService {
	return urlServiceImp{
		urlRepo:  urlRepo,
		userRepo: userRepo,
	}
}

func (urlService urlServiceImp) FindAll() ([]model.Url, error) {
	urls, err := urlService.urlRepo.GetAllUrl()
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func (urlService urlServiceImp) FindOne(userId, urlId int) (*model.Url, error) {
	user, err := urlService.urlRepo.GetUrlById(userId, urlId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (urlService urlServiceImp) Create(createUrlRequest dto.CreateUrlRequest) error {
	panic("url service imp")
}

func (urlService urlServiceImp) Delete(userId, urlId int) error {

	url, err := urlService.urlRepo.GetUrlById(userId, urlId)
	if err != nil {
		return err
	}
	user, err := urlService.userRepo.GetById(userId)
	if err != nil {
		return err
	}
	if user.ID != userId {
		return errors.New("something wrong")
	}
	return urlService.urlRepo.Delete(url.ID, user.ID)
}
