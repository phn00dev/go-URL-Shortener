package service

import (
	"errors"
	"log"

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

func (urlService urlServiceImp) FindAll() ([]model.Url, error) {
	urls, err := urlService.urlRepo.GetAllUrl()
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func (urlService urlServiceImp) FindOne(urlId int) (*model.Url, error) {
	user, err := urlService.urlRepo.GetUrlById(urlId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (urlService urlServiceImp) Create(userId int, createUrlRequest dto.CreateUrlRequest) (*model.Url, error) {

	// find user
	user, err := urlService.userRepo.GetById(userId)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("user not found")
	}
	newUrl := model.Url{
		OriginalUrl: createUrlRequest.OriginalUrl,
		ShortUrl:    generateshorturl.GenerateShortUrl(10),
		UserID:      user.ID,
		ClickCount:  0,
	}

	// Create new URL in DB
	if err := urlService.urlRepo.Create(newUrl); err != nil {
		return nil, err
	}

	// After creation, return the created URL object
	return &newUrl, nil
}

func (urlService urlServiceImp) Delete(userId, urlId int) error {
	// Ilki bilen URL-i tapmak
	url, err := urlService.urlRepo.GetUrlById(urlId)
	if err != nil {
		return err
	}

	// URL-iň userId-nin dogrylygyny barlamak
	if url.UserID != userId {
		return errors.New("user not authorized to delete this url")
	}

	// URL pozulýan bolsa, ony pozmak
	err = urlService.urlRepo.Delete(userId, urlId)
	if err != nil {
		return errors.New("failed to delete the URL")
	}

	return nil
}

// user urls

func (urlService urlServiceImp) FindAllUserUrls(userId int) ([]model.Url, error) {
	// get user
	user, err := urlService.userRepo.GetById(userId)
	if err != nil {
		return nil, err
	}
	// user urls
	userUrls, err := urlService.urlRepo.GetAllUserUrl(user.ID)
	if err != nil {
		return nil, err
	}
	return userUrls, nil
}

func (urlService urlServiceImp) FindOneUserUrl(userId, urlId int) (*model.Url, error) {
	user, err := urlService.userRepo.GetById(userId)
	if err != nil {
		return nil, err
	}
	url, err := urlService.urlRepo.GetOneUserUrl(user.ID, urlId)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (urlService urlServiceImp) GetByShortUrl(shortUrl string) (*model.Url, error) {

	if shortUrl == "" {
		return nil, errors.New("not empty url")
	}

	url, err := urlService.urlRepo.GetUrlByShortUrl(shortUrl)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (urlService urlServiceImp) UpdateClickCount(urlId, clickCount int) error {
	url, err := urlService.urlRepo.GetUrlById(urlId)
	if err != nil {
		return err
	}
	if url.ID == 0 {
		return errors.New("something wrong")
	}
	// update click count
	return urlService.urlRepo.UpdateUrlClickCount(url.ID, clickCount)
}

func (urlService urlServiceImp) SaveUrlAccessLog(accessLog model.UrlAccessLog) error {
	log.Println("service icinde", accessLog)
	return urlService.urlRepo.SaveUrlAccessLog(accessLog)
}
