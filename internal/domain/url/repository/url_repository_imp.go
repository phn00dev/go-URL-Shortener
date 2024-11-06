package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/phn00dev/go-URL-Shortener/internal/model"
)

type urlRepositoryImp struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return urlRepositoryImp{
		db: db,
	}
}

func (urlRepo urlRepositoryImp) GetUrlById(urlId int) (*model.Url, error) {
	var url model.Url
	if err := urlRepo.db.Where("id=?", urlId).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

func (urlRepo urlRepositoryImp) GetAllUrl() ([]model.Url, error) {
	var urls []model.Url
	if err := urlRepo.db.Order("id desc").Find(&urls).Error; err != nil {
		return nil, err
	}
	return urls, nil
}

func (urlRepo urlRepositoryImp) Create(url model.Url) error {
	return urlRepo.db.Create(&url).Error
}

func (urlRepo urlRepositoryImp) Delete(userId, urlId int) error {
	var url model.Url
	// Ilki bilen URL-äni gözlemek
	result := urlRepo.db.Where("user_id = ? AND id = ?", userId, urlId).Delete(&url)

	// Pozulýan maglumatyň barlygyny we hatalary barlamak
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows deleted")
	}
	return nil
}

func (urlRepo urlRepositoryImp) GetUrlShortUrl(shortUrl string) (*model.Url, error) {
	var url model.Url
	if err := urlRepo.db.Where("short_url=?", shortUrl).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

// user urls
func (urlRepo urlRepositoryImp) GetAllUserUrl(userId int) ([]model.Url, error) {
	var userUrls []model.Url
	if err := urlRepo.db.Where("user_id=?", userId).Find(&userUrls).Error; err != nil {
		return nil, err
	}
	return userUrls, nil
}

func (urlRepo urlRepositoryImp) GetOneUserUrl(userId, urlId int) (*model.Url, error) {
	var url model.Url
	if err := urlRepo.db.Where("user_id=? AND id=?", userId, urlId).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}
