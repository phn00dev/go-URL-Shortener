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
		ShortUrl:    generateshorturl.GenerateShortUrl(10),
		UserID:      user.ID,
		ClickCount:  0,
	}
	return urlService.urlRepo.Create(newUrl)
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
