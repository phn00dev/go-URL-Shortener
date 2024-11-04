package constructor

import (
	"gorm.io/gorm"

	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/handler"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/repository"
	"github.com/phn00dev/go-URL-Shortener/internal/domain/admin/service"
)

var (
	adminRepo    repository.AdminRepository
	adminService service.AdminService
	AdminHandler handler.AdminHandler
)

func InitAdminRequirements(db *gorm.DB) {

	adminRepo = repository.NewAdminRepository(db)
	adminService = service.NewAdminService(adminRepo)
	AdminHandler = handler.NewAdminHandler(adminService)
}
