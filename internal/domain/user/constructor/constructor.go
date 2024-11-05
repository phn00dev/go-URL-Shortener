package constructor

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/handler"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/repository"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/user/service"
)

var (
	userRepo    repository.UserRepository
	userService service.UserService
	UserHandler handler.UserHandler
)

func InitUserRequirements(db *gorm.DB) {
	userRepo = repository.NewUserRepository(db)
	userService = service.NewUserService(userRepo)
	UserHandler = handler.NewUserHandler(userService)
}
