package service

import (
	"errors"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/dto"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/repository"
	userRepository "github.com/phn00dev/go-URL-Shortener/internal/domain/user/repository"
	"github.com/phn00dev/go-URL-Shortener/internal/model"
	generateshorturl "github.com/phn00dev/go-URL-Shortener/internal/utils/generate_short_url"
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

func (urlService urlServiceImp) FindAll(userId int) ([]model.Url, error) {
	urls, err := urlService.urlRepo.GetAllUrl(userId)
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

func (urlService urlServiceImp) Create(userId int, createUrlRequest dto.CreateUrlRequest) error {

	// find user
	user, err := urlService.userRepo.GetById(userId)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("user not found")
	}
	newUrl := model.Url{
		OriginalUrl: createUrlRequest.OriginalUrl,
		ShortUrl:    generateshorturl.GenerateShortUrl(11),
		UserID:      user.ID,
		ClickCount:  0,
	}
	return urlService.urlRepo.Create(newUrl)
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
