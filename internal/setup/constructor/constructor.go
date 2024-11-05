package constructor

import (
	"github.com/phn00dev/go-URL-Shortener/internal/app"
	adminConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/admin/constructor"
	urlConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/url/constructor"
	userConstructor "github.com/phn00dev/go-URL-Shortener/internal/domain/user/constructor"
)

func InitDependencies(dependencies *app.Dependencies) {
	adminConstructor.InitAdminRequirements(dependencies.DB)
	userConstructor.InitUserRequirements(dependencies.DB)
	urlConstructor.InitUrlRequirements(dependencies.DB)
}
