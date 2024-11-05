package constructor

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/handler"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/repository"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/url/service"
	userRepository "github.com/phn00dev/go-URL-Shortener/internal/domain/user/repository"
)

var (
	urlRepo    repository.UrlRepository
	userRepo   userRepository.UserRepository
	urlService service.UrlService
	UrlHandler handler.UrlHandler
)

func InitUrlRequirements(db *gorm.DB) {
	urlRepo = repository.NewUrlRepository(db)
	userRepo = userRepository.NewUserRepository(db)
	urlService = service.NewUrlService(urlRepo, userRepo)
	UrlHandler = handler.NewUrlHandler(urlService)
}
